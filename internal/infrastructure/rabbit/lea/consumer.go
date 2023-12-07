package leawash

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"sbp/internal/app"
	"sbp/internal/conversions"
	"sbp/internal/entities"
	rabbitEntities "sbp/internal/entities/rabbit"
	"sbp/pkg/rabbitmq"

	gorabbit "github.com/wagslane/go-rabbitmq"

	"go.uber.org/zap"
)

type leaClient struct {
	logger    *zap.SugaredLogger
	consumer  *rabbitmq.Consumer
	publisher app.LeaWashPublisher
}

type washHandler interface {
	InitPayment(ctx context.Context, payRequest entities.PaymentRequest) (payResponse *entities.PaymentResponse, err error)
	CancelPayment(ctx context.Context, req entities.PaymentСancellationRequest) (resendNeaded bool, err error)
}

func NewLeaConsumer(logger *zap.SugaredLogger, client *rabbitmq.RabbitMqClient, washHandler washHandler, publisher app.LeaWashPublisher) (*leaClient, error) {
	if client == nil {
		return nil, errors.New("NewLeaConsumer: client = nil")
	}

	exchangeName := rabbitEntities.SbpClientExchange
	routingKey := rabbitEntities.SbpClientRoutingKey

	handler, err := createHandler(logger, washHandler, publisher)
	if err != nil {
		return nil, err
	}

	consumer, err := rabbitmq.NewConsumer(
		logger,
		client,
		handler,
		string(routingKey),
		gorabbit.WithConsumerOptionsExchangeDeclare,

		gorabbit.WithConsumerOptionsExchangeName(string(exchangeName)),
		gorabbit.WithConsumerOptionsExchangeKind("direct"),

		gorabbit.WithConsumerOptionsRoutingKey(string(routingKey)),
		gorabbit.WithConsumerOptionsExchangeDurable,
		gorabbit.WithConsumerOptionsConsumerExclusive,
	)
	if err != nil {
		return nil, err
	}

	return &leaClient{
		logger:    logger,
		consumer:  consumer,
		publisher: publisher,
	}, nil
}

func (c leaClient) Close() {
	c.consumer.Close()
}

func createHandler(logger *zap.SugaredLogger, washHandler washHandler, publisher app.LeaWashPublisher) (rabbitmq.ConsumerHandler, error) {
	return func(ctx context.Context, d rabbitmq.RbqMessage) error {
		messageType := rabbitEntities.Message(d.Type)
		switch messageType {
		case rabbitEntities.PaymentRequestMessage:
			{
				var req rabbitEntities.PaymentRequest
				err := json.Unmarshal(d.Body, &req)
				if err != nil {
					logger.Debugf("lea payment request error: '%s'", err.Error())
					return err
				}

				logger.Debugf("lea payment request from wash_id: '%s', post_id: '%s'", req.WashID, req.PostID)

				sbpRequest := conversions.PaymentRequestToSbp(req)
				payResp, err := washHandler.InitPayment(ctx, sbpRequest)
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
		case rabbitEntities.PaymentСancellationRequestMessage:
			{
				var req rabbitEntities.PaymentСancellationRequest
				err := json.Unmarshal(d.Body, &req)
				if err != nil {
					logger.Debugf("lea payment cancellation request error: '%s'", err.Error())
					return err
				}

				logger.Debugf("lea payment cancellation request from wash_id: '%s', post_id: '%s'", req.WashID, req.PostID)

				sbpRequest := conversions.PaymentСancellationRequestToSbp(req)
				resendNeaded, err := washHandler.CancelPayment(ctx, sbpRequest)
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
