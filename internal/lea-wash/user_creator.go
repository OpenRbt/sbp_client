package leawash

import (
	logicEntities "sbp/internal/logic/entities"
	rabbitMq "sbp/internal/rabbit-mq"
)

// brokerUserCreator ...
type brokerUserCreator struct {
	rabbitMqClient *rabbitMq.RabbitMqClient
}

// NewBrokerUserCreator ...
func NewBrokerUserCreator(rabbitMqClient *rabbitMq.RabbitMqClient) (*brokerUserCreator, error) {
	return &brokerUserCreator{
		rabbitMqClient: rabbitMqClient,
	}, nil
}

// CreateUser ...
func (c *brokerUserCreator) CreateUser(login string, password string) error {
	return c.rabbitMqClient.CreateUser(string(logicEntities.ServiceLeaCentralWash), login, password)
}
