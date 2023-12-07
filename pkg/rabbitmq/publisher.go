package rabbitmq

import (
	"encoding/json"
	"fmt"
	rabbitEntities "sbp/internal/entities/rabbit"

	rabbitmq "github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

type Publisher struct {
	logger          *zap.SugaredLogger
	eventsPublisher *rabbitmq.Publisher
	exchangeName    string
}

func NewPublisher(
	logger *zap.SugaredLogger,
	rabbitMqClient *RabbitMqClient,
	exchangeName string,
) (*Publisher, error) {

	if rabbitMqClient == nil {
		return nil, fmt.Errorf("NewPublisher: rabbitMqClient == nil")
	}

	newPublisher, err := rabbitMqClient.NewPublisher(exchangeName)
	if err != nil {
		return nil, fmt.Errorf("NewPublisher: %s", err)
	}

	return &Publisher{
		logger:          logger,
		eventsPublisher: newPublisher,
		exchangeName:    exchangeName,
	}, nil
}

func (p *Publisher) Send(
	messageStruct interface{},
	service rabbitEntities.Exchange,
	routingKey rabbitEntities.RoutingKey,
	messageType rabbitEntities.Message,
) error {

	if messageStruct == nil {
		return fmt.Errorf("wash message is nil")
	}

	m, err := json.Marshal(&messageStruct)
	if err != nil {
		return err
	}

	exchangeName := string(service)
	err = p.eventsPublisher.Publish(
		m,
		[]string{string(routingKey)},
		rabbitmq.WithPublishOptionsType(string(messageType)),
		rabbitmq.WithPublishOptionsExchange(exchangeName),
	)

	return err
}
