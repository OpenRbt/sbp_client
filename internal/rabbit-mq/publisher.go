package rabbitmq

import (
	"encoding/json"
	"fmt"
	rabbitmqClient "sbp/pkg/rabbit-mq"

	shareBusinessEntities "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	rabbitmq "github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

// Publisher ...
type Publisher struct {
	logger          *zap.SugaredLogger
	eventsPublisher *rabbitmq.Publisher
	exchangeName    string
}

// newPublisher ...
func newPublisher(
	logger *zap.SugaredLogger,
	rabbitMqClient *rabbitmqClient.RabbitMqClient,
	exchangeName string,
) (*Publisher, error) {

	if rabbitMqClient == nil {
		return nil, fmt.Errorf("NewPublisher: rabbitMqClient == nil")
	}

	// publisher
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

// Send ...
func (p *Publisher) Send(
	messageStruct interface{},
	service shareBusinessEntities.Service,
	routingKey shareBusinessEntities.RoutingKey,
	messageType shareBusinessEntities.MessageType,
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
