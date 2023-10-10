package leawash

import (
	leEntities "sbp/internal/lea-wash/entities"
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
	return c.rabbitMqClient.CreateUser(
		string(leEntities.ExchangeNameLeaCentralWash),
		string(leEntities.ExchangeNameSbpClient),
		login,
		password)
}
