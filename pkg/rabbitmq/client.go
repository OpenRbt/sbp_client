package rabbitmq

import (
	"context"
	"fmt"
	"sbp/internal/config"
	rabbitmqGeneratedClient "sbp/rabbitmqapi/client"
	rabbitmqGeneratedOperations "sbp/rabbitmqapi/client/operations"
	rabbitmqGenerateEntities "sbp/rabbitmqapi/models"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	amqp "github.com/rabbitmq/amqp091-go"
	rabbitmq "github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

type ServerID string

type RabbitMqClient struct {
	logger     *zap.SugaredLogger
	connection *rabbitmq.Conn
	httpApi    *rabbitmqGeneratedClient.RabbitIntAPI
	httpAuth   runtime.ClientAuthInfoWriter
}

func NewRabbitMqClient(config config.RabbitMqClientConfig) (*RabbitMqClient, error) {
	rabbitUser := config.RabbitMQConfig.User
	rabbitPassword := config.RabbitMQConfig.Password
	rabbitUrl := config.RabbitMQConfig.Url
	rabbitPort := config.RabbitMQConfig.Port

	connString := ""
	if config.RabbitMQConfig.Secure {
		connString = fmt.Sprintf("amqps://%s:%s@%s:%s/",
			rabbitUser,
			rabbitPassword,
			rabbitUrl,
			rabbitPort,
		)
	} else {
		connString = fmt.Sprintf("amqp://%s:%s@%s:%s/",
			rabbitUser,
			rabbitPassword,
			rabbitUrl,
			rabbitPort,
		)
	}

	rabbitConf := rabbitmq.Config{
		SASL: []amqp.Authentication{
			&amqp.PlainAuth{
				Username: rabbitUser,
				Password: rabbitPassword,
			},
		},
		Vhost:      "/",
		ChannelMax: 0,
		FrameSize:  0,
		Heartbeat:  0,
		Properties: nil,
		Locale:     "",
		Dial:       nil,
	}

	conn, err := rabbitmq.NewConn(
		connString,
		rabbitmq.WithConnectionOptionsLogging,
		rabbitmq.WithConnectionOptionsConfig(rabbitConf),
	)
	if err != nil {
		return nil, err
	}

	host := fmt.Sprintf("%s:%s", rabbitUrl, config.RabbitMQConfig.PortWeb)
	path := ""
	shemes := []string{"http"}
	transport := httptransport.New(host, path, shemes)
	httpApi := rabbitmqGeneratedClient.New(transport, strfmt.Default)
	httpAuth := httptransport.BasicAuth(rabbitUser, rabbitPassword)

	return &RabbitMqClient{
		connection: conn,
		logger:     config.Logger,
		httpApi:    httpApi,
		httpAuth:   httpAuth,
	}, nil
}

func (c *RabbitMqClient) Connection() *rabbitmq.Conn {
	return c.connection
}

func (c *RabbitMqClient) NewPublisher(exchangeName string) (*rabbitmq.Publisher, error) {
	publisher, err := rabbitmq.NewPublisher(
		c.connection,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsExchangeKind("direct"),
		rabbitmq.WithPublisherOptionsExchangeName(exchangeName),
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	return publisher, err
}

// func (c *RabbitMqClient) NewConsumer(exchangeName string, routingKey string, queue string, handler func(d rabbitmq.Delivery) (action rabbitmq.Action)) (*rabbitmq.Consumer, error) {
// 	return rabbitmq.NewConsumer(
// 		c.connection,
// 		handler,
// 		queue,
		// rabbitmq.WithConsumerOptionsExchangeDeclare,

		// rabbitmq.WithConsumerOptionsExchangeName(exchangeName),
		// rabbitmq.WithConsumerOptionsExchangeKind("direct"),

		// rabbitmq.WithConsumerOptionsRoutingKey(routingKey),
		// rabbitmq.WithConsumerOptionsExchangeDurable,
		// rabbitmq.WithConsumerOptionsConsumerExclusive,
// 	)
// }

type RabbitHandler = func(d rabbitmq.Delivery) (action rabbitmq.Action)

func (c *RabbitMqClient) NewConsumer(handler RabbitHandler, queue string, args ...func(*rabbitmq.ConsumerOptions)) (*rabbitmq.Consumer, error) {
	return rabbitmq.NewConsumer(
		c.connection,
		handler,
		queue,
		args...,
	)
}

func (c *RabbitMqClient) CreateExchangeName(exchangeName string) error {
	ctx := context.TODO()

	createNewExchangeParams := &rabbitmqGeneratedOperations.CreateNewExchangeParams{
		Body: &rabbitmqGenerateEntities.ExchangeProperties{
			AutoDelete: false,
			Durable:    true,
			Internal:   false,
			Type:       "direct",
		},
		ExchangeName: exchangeName,
		Vhost:        "/",
		Context:      ctx,
		HTTPClient:   nil,
	}

	_, _, err := c.httpApi.Operations.CreateNewExchange(createNewExchangeParams, c.httpAuth)
	if err != nil {
		return err
	}

	return nil
}

func (c *RabbitMqClient) CreateRabbitUser(readExchangeName string, writeExchangeName string, login string, password string) (err error) {
	ctx := context.TODO()

	tags := ""
	vhost := "/"

	createUserParams := &rabbitmqGeneratedOperations.CreateUserParams{
		Body: &rabbitmqGenerateEntities.CreateUser{
			Password: &password,
			Tags:     &tags,
		},
		UserID:     login,
		Context:    ctx,
		HTTPClient: nil,
	}
	_, _, err = c.httpApi.Operations.CreateUser(createUserParams, c.httpAuth)

	if err != nil {
		return err
	}

	setUserPermsParams := &rabbitmqGeneratedOperations.SetUserPermsParams{
		Body: &rabbitmqGenerateEntities.ManagePermissions{
			Configure: fmt.Sprintf("%s.*", login),
			Read:      fmt.Sprintf("(%s)|(%s).*", readExchangeName, login),
			Write:     fmt.Sprintf("(%s)|(%s).*", writeExchangeName, login),
		},
		UserID:     login,
		Vhost:      vhost,
		Context:    ctx,
		HTTPClient: nil,
	}
	_, _, err = c.httpApi.Operations.SetUserPerms(setUserPermsParams, c.httpAuth)
	if err != nil {
		return err
	}

	return err
}
