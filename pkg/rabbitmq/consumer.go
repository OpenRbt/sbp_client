package rabbitmq

import (
	"context"
	"fmt"

	rabbitmq "github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

type Consumer struct {
	logger         *zap.SugaredLogger
	eventsConsumer *rabbitmq.Consumer
}

type RbqMessage rabbitmq.Delivery

type ConsumerHandler func(ctx context.Context, d RbqMessage) error

type ExchangeKind string

const (
	DirectExchangeKind ExchangeKind = "direct"
	FanoutExchangeKind ExchangeKind = "fanout"
)

func NewConsumer(
	logger *zap.SugaredLogger,
	client *RabbitMqClient,
	handler ConsumerHandler,
	queue string,
	args ...func(*rabbitmq.ConsumerOptions),
) (*Consumer, error) {
	consumerHandler := func(d rabbitmq.Delivery) (action rabbitmq.Action) {
		ctx := context.TODO()
		err := handler(ctx, RbqMessage(d))
		if err != nil {
			logger.Error(err)
			return rabbitmq.NackDiscard
		}

		return rabbitmq.Ack
	}

	newConsumer, err := client.NewConsumer(consumerHandler, queue, args...)
	if err != nil {
		return nil, fmt.Errorf("NewConsumer: %s", err)
	}

	return &Consumer{
		logger:         logger,
		eventsConsumer: newConsumer,
	}, nil
}

func (c *Consumer) Close() {
	c.eventsConsumer.Close()
}
