package rabbitmq

import (
	"context"
	"fmt"
	rabbitmqClient "sbp/pkg/rabbit-mq"

	rabbitmq "github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

// Consumer ...
type Consumer struct {
	logger         *zap.SugaredLogger
	eventsConsumer *rabbitmq.Consumer
}

// RbqMessage ...
type RbqMessage rabbitmq.Delivery

// ConsumerHandler ...
type ConsumerHandler func(ctx context.Context, d RbqMessage) error

// newConsumer ...
func newConsumer(
	logger *zap.SugaredLogger,
	client *rabbitmqClient.RabbitMqClient,
	handler ConsumerHandler,
	exchangeName string,
	routingKey string,
) (*Consumer, error) {
	// consumer handler
	consumerHandler := func(d rabbitmq.Delivery) (action rabbitmq.Action) {
		ctx := context.TODO()
		err := handler(ctx, RbqMessage(d))
		if err != nil {
			logger.Error(err)
			return rabbitmq.NackDiscard
		}
		return rabbitmq.Ack
	}

	newConsumer, err := client.NewConsumer(exchangeName, routingKey, consumerHandler)
	if err != nil {
		return nil, fmt.Errorf("NewConsumer: %s", err)
	}

	return &Consumer{
		logger:         logger,
		eventsConsumer: newConsumer,
	}, nil
}

// Close ...
func (c *Consumer) Close() {
	c.eventsConsumer.Close()
}
