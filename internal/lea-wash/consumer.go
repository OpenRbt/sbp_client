package leawash

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	leEntities "sbp/internal/lea-wash/entities"

	"sbp/internal/logic"
	rabbitMq "sbp/internal/rabbit-mq"

	shareBusinessEntities "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"

	leConverter "sbp/internal/lea-wash/converter"
	logicEntities "sbp/internal/logic/entities"

	"go.uber.org/zap"
)

// leaWashConsumer ...
type leaWashConsumer struct {
	logger    *zap.SugaredLogger
	consumer  *rabbitMq.Consumer
	publisher logic.LeaWashPublisher
}

// washHandler ...
type washHandler interface {
	Pay(ctx context.Context, payRequest logicEntities.PaymentRequest) (payResponse *logicEntities.PaymentResponse, err error)
	Cancel(ctx context.Context, req logicEntities.PaymentСancellationRequest) (resendNeaded bool, err error)
}

// NewLeaWashConsumer ...
func NewLeaWashConsumer(logger *zap.SugaredLogger, client *rabbitMq.RabbitMqClient, washHandler washHandler, publisher logic.LeaWashPublisher) (*leaWashConsumer, error) {
	if client == nil {
		return nil, errors.New("NewLeaWashConsumer: client = nil")
	}

	exchangeName := leEntities.ExchangeNameSbpClient
	routingKey := leEntities.RoutingKeySbpClient

	handler, err := createHandler(logger, washHandler, publisher)
	if err != nil {
		return nil, err
	}

	consumer, err := client.CreateConsumer(string(exchangeName), string(routingKey), handler)
	if err != nil {
		return nil, err
	}

	return &leaWashConsumer{
		logger:    logger,
		consumer:  consumer,
		publisher: publisher,
	}, nil
}

// Close ...
func (c leaWashConsumer) Close() {
	c.consumer.Close()
}

// processMessage ...
func createHandler(logger *zap.SugaredLogger, washHandler washHandler, publisher logic.LeaWashPublisher) (rabbitMq.ConsumerHandler, error) {
	return func(ctx context.Context, d rabbitMq.RbqMessage) error {
		messageType := shareBusinessEntities.MessageType(d.Type)
		switch messageType {
		case leEntities.MessageTypePaymentRequest:
			{
				var req leEntities.PaymentRequest
				err := json.Unmarshal(d.Body, &req)
				if err != nil {
					logger.Debugf("lea payment request error: '%s'", err.Error())
					return err
				}

				logger.Debugf("lea payment request from wash_id: '%s', post_id: '%s'", req.WashID, req.PostID)

				sbpRequest := leConverter.PaymentRequestToSbp(req)
				payResp, err := washHandler.Pay(ctx, sbpRequest)
				if err != nil {
					logger.Debugf("lea payment request error: '%s'", err.Error())
					errPub := publisher.SendToLeaPaymentFailedResponse(sbpRequest.WashID, sbpRequest.PostID, sbpRequest.OrderID, err.Error())
					return errPub
				}
				if payResp == nil {
					logger.Debug("lea payment request error: 'payment_resp = nil'")
					logger.Debugf("lea payment request error: '%s'", err.Error())
					err = publisher.SendToLeaPaymentFailedResponse(sbpRequest.WashID, sbpRequest.PostID, sbpRequest.OrderID, err.Error())
					return err
				}

				return nil

			}
		case leEntities.MessageTypePaymentСancellationRequest:
			{
				var req leEntities.PaymentСancellationRequest
				err := json.Unmarshal(d.Body, &req)
				if err != nil {
					logger.Debugf("lea payment cancellation request error: '%s'", err.Error())
					return err
				}

				logger.Debugf("lea payment cancellation request from wash_id: '%s', post_id: '%s'", req.WashID, req.PostID)

				sbpRequest := leConverter.PaymentСancellationRequestToSbp(req)
				resendNeaded, err := washHandler.Cancel(ctx, sbpRequest)
				if err != nil {
					logger.Debugf("lea payment cancellation request error: '%s'", err.Error())
					if resendNeaded {
						err = publisher.SendToLeaPaymentFailedResponse(sbpRequest.WashID, sbpRequest.PostID, sbpRequest.OrderID, err.Error())
						return err
					}
					return err
				}

				return nil
			}
		default:
			return fmt.Errorf("received unexpected message with type: %s", messageType)
		}
	}, nil
}
