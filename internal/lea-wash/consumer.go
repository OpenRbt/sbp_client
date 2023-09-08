package leawash

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sbp/internal/logic"
	logicEntities "sbp/internal/logic/entities"
	rabbitMq "sbp/internal/rabbit-mq"

	shareBusinessEntities "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
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
	Pay(ctx context.Context, payRequest logicEntities.PayRequest) (payResponse *logicEntities.PayResponse, err error)
	Cancel(ctx context.Context, req logicEntities.PayСancellationRequest) (resendNeaded bool, err error)
}

// NewLeaWashConsumer ...
func NewLeaWashConsumer(logger *zap.SugaredLogger, client *rabbitMq.RabbitMqClient, washHandler washHandler, publisher logic.LeaWashPublisher) (*leaWashConsumer, error) {
	if client == nil {
		return nil, errors.New("NewLeaWashConsumer: client = nil")
	}

	exchangeName := logicEntities.ServiceSbpClient
	routingKey := logicEntities.RoutingKeySbpClient

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
		case logicEntities.MessageTypePaymentRequest:
			{
				var req logicEntities.PayRequest
				err := json.Unmarshal(d.Body, &req)
				if err != nil {
					return err
				}

				payResp, err := washHandler.Pay(ctx, req)
				if err != nil {
					err = publisher.SendToLeaError(req.WashID, req.PostID, req.OrderID, err.Error(), logicEntities.ErrorPaymentRequestFailed)
					return err
				}

				if payResp != nil {
					return publisher.SendToLea(req.WashID, string(logicEntities.MessageTypePaymentResponse), payResp)
				}

				return errors.New("payment_resp = nil")
			}
		case logicEntities.MessageTypePaymentСancellationRequest:
			{
				var req logicEntities.PayСancellationRequest
				err := json.Unmarshal(d.Body, &req)
				if err != nil {
					return err
				}

				resendNeaded, err := washHandler.Cancel(ctx, req)
				if err != nil {
					if resendNeaded {
						err = publisher.SendToLeaError(req.WashID, req.PostID, req.OrderID, err.Error(), logicEntities.ErrorPaymentСancellationFailed)
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
