package leawash

import (
	rabbitEntities "sbp/internal/entities/rabbit"
	"sbp/pkg/rabbitmq"
)

type brokerUserCreator struct {
	rabbitMqClient *rabbitmq.RabbitMqClient
}

func NewBrokerUserCreator(rabbitMqClient *rabbitmq.RabbitMqClient) (*brokerUserCreator, error) {
	return &brokerUserCreator{
		rabbitMqClient: rabbitMqClient,
	}, nil
}

func (c *brokerUserCreator) CreateUser(login string, password string) error {
	return c.rabbitMqClient.CreateRabbitUser(
		string(rabbitEntities.LeaCentralWashExchange),
		string(rabbitEntities.SbpClientExchange),
		login,
		password)
}
