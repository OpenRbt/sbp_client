package shareRabbit

import (
	"errors"
	"sbp/internal/app"
	rabbitEntities "sbp/internal/entities/rabbit"
	"sbp/pkg/rabbitmq"

	gorabbit "github.com/wagslane/go-rabbitmq"

	"go.uber.org/zap"
)

var _ = app.SharePublisher(&sharePublisher{})

type sharePublisher struct {
	*gorabbit.Publisher
}

func NewSharePublisher(logger *zap.SugaredLogger, rabbitMqClient *rabbitmq.RabbitMqClient) (*sharePublisher, error) {
	if rabbitMqClient == nil {
		return nil, errors.New("unable to create share publisher: rabbit client = nil")
	}

	p, err := gorabbit.NewPublisher(
		rabbitMqClient.Connection(),
		gorabbit.WithPublisherOptionsLogging,
		gorabbit.WithPublisherOptionsExchangeDeclare,
		gorabbit.WithPublisherOptionsExchangeName(string(rabbitEntities.WashBonusExchange)),
		gorabbit.WithPublisherOptionsExchangeKind("direct"),
		gorabbit.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	return &sharePublisher{p}, nil
}

func (pub *sharePublisher) SendDataRequest() error {
	return pub.Publish(
		nil,
		[]string{string(rabbitEntities.WashBonusRoutingKey)},
		gorabbit.WithPublishOptionsType(string(rabbitEntities.RequestAdminDataMessage)),
		gorabbit.WithPublishOptionsReplyTo(string(rabbitEntities.SBPStartupQueue)),
	)
}
