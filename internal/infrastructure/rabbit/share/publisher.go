package shareRabbit

import (
	"encoding/json"
	"errors"
	"fmt"
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
		gorabbit.WithPublishOptionsType(string(rabbitEntities.RequestAdminDataMessageType)),
		gorabbit.WithPublishOptionsReplyTo(string(rabbitEntities.SBPStartupQueue)),
	)
}

func (pub *sharePublisher) CreateRabbitUser(login, password string) error {
	return pub.send(
		rabbitEntities.UserCreation{
			ID:            login,
			ServiceKey:    password,
			ReadExchange:  string(rabbitEntities.LeaCentralWashExchange),
			WriteExchange: string(rabbitEntities.SbpClientExchange),
		},
		rabbitEntities.WashBonusExchange,
		rabbitEntities.WashBonusRoutingKey,
		rabbitEntities.CreateUserMessageType,
	)
}

func (pub *sharePublisher) send(message interface{}, exchange rabbitEntities.Exchange, routingKey rabbitEntities.RoutingKey, messageType rabbitEntities.MessageType) error {
	if message == nil {
		return fmt.Errorf("wash message is nil")
	}

	byteMessage, err := json.Marshal(&message)
	if err != nil {
		return err
	}

	exchangeName := string(exchange)
	err = pub.Publish(
		byteMessage,
		[]string{string(routingKey)},
		gorabbit.WithPublishOptionsType(string(messageType)),
		gorabbit.WithPublishOptionsExchange(exchangeName),
	)

	return err
}
