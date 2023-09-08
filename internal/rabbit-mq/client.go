package rabbitmq

import (
	"errors"

	// generated

	"sbp/pkg/bootstrap"

	// go-openapi

	// zab
	"go.uber.org/zap"

	// rabbitmq client
	rabbitmqClient "sbp/pkg/rabbit-mq"
)

const layer = "lea-cw-client"

// ServerID...
type ServerID string

// RabbitMqClient ...
type RabbitMqClient struct {
	logger         *zap.SugaredLogger
	rabbitMqClient *rabbitmqClient.RabbitMqClient
}

// RabbitMqClientConfig ...
type RabbitMqClientConfig struct {
	Logger         *zap.SugaredLogger
	RabbitMQConfig *bootstrap.RabbitMQConfig
}

// checkRabbitMqClientConfig ...
func checkRabbitMqClientConfig(conf RabbitMqClientConfig) error {
	if conf.Logger == nil {
		return errors.New("logger is empty")
	}
	if conf.RabbitMQConfig == nil {
		return errors.New("rabbit_mq_config is empty")
	}
	return nil
}

// NewRabbitMqClient ...
func NewRabbitMqClient(config RabbitMqClientConfig) (*RabbitMqClient, error) {
	// check config
	err := checkRabbitMqClientConfig(config)
	if err != nil {
		return nil, bootstrap.CustomError(layer, "checkRabbitMqClientConfig", err)
	}

	newRabbitMqClient, err := rabbitmqClient.NewRabbitMqClient(*config.RabbitMQConfig, config.Logger)
	if err != nil {
		return nil, bootstrap.CustomError(layer, "NewRabbitMqClient", err)
	}

	c := &RabbitMqClient{
		logger:         config.Logger,
		rabbitMqClient: newRabbitMqClient,
	}

	return c, nil
}

// CreateConsumer ...
func (c *RabbitMqClient) CreateConsumer(exchangeName string, routingKey string, handler ConsumerHandler) (*Consumer, error) {
	return newConsumer(c.logger, c.rabbitMqClient, handler, exchangeName, routingKey)
}

// CreatePublisher ...
func (c *RabbitMqClient) CreatePublisher(exchangeName string) (*Publisher, error) {
	return newPublisher(c.logger, c.rabbitMqClient, exchangeName)
}

// CreateUser ...
func (c *RabbitMqClient) CreateUser(exchangeName string, login string, password string) error {
	return c.rabbitMqClient.CreateRabbitUser(exchangeName, login, password)
}
