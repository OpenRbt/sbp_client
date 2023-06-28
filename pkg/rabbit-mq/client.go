package rabbitmq

import (
	"context"
	"fmt"
	"sbp/pkg/bootstrap"

	rabbitmqGeneratedClient "sbp/pkg/rabbit-mq/client"
	rabbitmqGeneratedOperations "sbp/pkg/rabbit-mq/client/operations"
	rabbitmqGenerateEntities "sbp/pkg/rabbit-mq/models"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

// RabbitMqClient ...
type RabbitMqClient struct {
	logger     *zap.SugaredLogger
	connection *rabbitmq.Conn
	httpApi    *rabbitmqGeneratedClient.RabbitIntAPI
	httpAuth   runtime.ClientAuthInfoWriter
}

// NewRabbitMqClient ...
func NewRabbitMqClient(config bootstrap.RabbitMQConfig, logger *zap.SugaredLogger) (*RabbitMqClient, error) {
	rabbitUser := config.User
	rabbitPassword := config.Password
	rabbitUrl := config.Url
	rabbitPort := config.Port

	connString := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		rabbitUser,
		rabbitPassword,
		rabbitUrl,
		rabbitPort,
	)

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

	// connection
	conn, err := rabbitmq.NewConn(
		connString,
		rabbitmq.WithConnectionOptionsLogging,
		rabbitmq.WithConnectionOptionsConfig(rabbitConf),
	)
	if err != nil {
		return nil, err
	}

	// http
	host := rabbitUrl + ":15672"
	path := ""
	shemes := []string{"http"}
	transport := httptransport.New(host, path, shemes)
	httpApi := rabbitmqGeneratedClient.New(transport, strfmt.Default)
	httpAuth := httptransport.BasicAuth(rabbitUser, rabbitPassword)

	return &RabbitMqClient{
		connection: conn,
		logger:     logger,
		httpApi:    httpApi,
		httpAuth:   httpAuth,
	}, nil
}

// NewConsumer ...
func (c *RabbitMqClient) Connection() *rabbitmq.Conn {
	return c.connection
}

// NewConsumer ...
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

// NewConsumer ...
func (c *RabbitMqClient) NewConsumer(exchangeName string, routingKey string, handler func(d rabbitmq.Delivery) (action rabbitmq.Action)) (*rabbitmq.Consumer, error) {

	consumer, err := rabbitmq.NewConsumer(
		c.connection,
		handler,
		routingKey,
		rabbitmq.WithConsumerOptionsExchangeDeclare,

		// SbpAdminService
		rabbitmq.WithConsumerOptionsExchangeName(exchangeName),
		rabbitmq.WithConsumerOptionsExchangeKind("direct"),

		// SbpAdminRoutingKey
		rabbitmq.WithConsumerOptionsRoutingKey(routingKey),
		rabbitmq.WithConsumerOptionsExchangeDurable,
		rabbitmq.WithConsumerOptionsConsumerExclusive,
	)

	return consumer, err
}

// CreateRabbitUser ...
func (c *RabbitMqClient) CreateRabbitUser(exchangeName string, userId string, userKey string) (err error) {
	ctx := context.TODO()

	tags := ""
	vhost := "/"

	// create user ...
	createUserParams := &rabbitmqGeneratedOperations.CreateUserParams{
		Body: &rabbitmqGenerateEntities.CreateUser{
			Password: &userKey,
			Tags:     &tags,
		},
		UserID:     userId,
		Context:    ctx,
		HTTPClient: nil,
	}
	_, _, err = c.httpApi.Operations.CreateUser(createUserParams, c.httpAuth)

	if err != nil {
		return err
	}

	// set user perms
	setUserPermsParams := &rabbitmqGeneratedOperations.SetUserPermsParams{
		Body: &rabbitmqGenerateEntities.ManagePermissions{
			Configure: fmt.Sprintf("%s.*", userId),
			Read:      fmt.Sprintf("(%s)|(%s).*", exchangeName, userId),
			Write:     fmt.Sprintf("(%s)|(%s).*", exchangeName, userId),
		},
		UserID:     userId,
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
