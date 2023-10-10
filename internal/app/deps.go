package app

import (
	"context"
	"time"

	restApi "sbp/internal/api/rest"
	authClient "sbp/internal/auth-client/firebase"
	leawash "sbp/internal/lea-wash"
	logic "sbp/internal/logic"
	payClient "sbp/internal/pay-client"
	rabbitmq "sbp/internal/rabbit-mq"
	repository "sbp/internal/repository"
	server "sbp/internal/server"
	timerTriggeredScheduler "sbp/internal/timer-triggered-scheduler"
	"sbp/pkg/bootstrap"

	"go.uber.org/zap"
)

// deps ...
type deps struct {
	// common
	config *bootstrap.Config
	logger *zap.SugaredLogger

	// clients
	repository        logic.Repository
	authClient        logic.AuthClient
	rabbitMqClient    *rabbitmq.RabbitMqClient
	payClient         logic.PayClient
	leaWashPublisher  logic.LeaWashPublisher
	leaWashConsumer   leaWashConsumer
	brokerUserCreator logic.BrokerUserCreator

	// logic
	logic *logic.Logic

	// api
	api server.Api

	// server
	httpServer HttpServer

	// timer triggered
	timerTriggeredScheduler timerTriggeredScheduler.TimerTriggeredScheduler

	//closeFuncs
	closeFuncs []func() error
}

type leaWashConsumer interface {
	Close()
}

func InitDeps(ctx context.Context, config *bootstrap.Config, logger *zap.SugaredLogger) (*deps, error) {
	d := deps{
		config: config,
		logger: logger,
	}

	inits := []func(context.Context) error{
		d.initRabbitMq,
		d.initRepository,
		d.initAuthClient,
		d.initPayClient,
		d.initLeaWashPublisher,
		d.initBrokerUserCreator,
		d.initLogic,
		d.initLeaWashConsumer,
		d.initHttpApi,
		d.initHttpServer,
		d.initTimerTriggeredScheduler,
	}

	d.logger.Info("start init deps")
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return nil, err
		}
	}

	return &d, nil
}

// close ...
func (d *deps) close() {
	closeFuncs := []func(){
		d.leaWashConsumer.Close,
	}
	d.logger.Info("start close deps")
	for _, f := range closeFuncs {
		f()
	}
}

// addCloser ...
func (d *deps) addCloser(f func() error) {
	d.closeFuncs = append(d.closeFuncs, f)
}

// initRepository ...
func (d *deps) initRepository(ctx context.Context) (err error) {
	repositoryConfig := repository.RepositoryConfig{
		DBConfig: &d.config.DB,
		Logger:   d.logger,
	}
	d.repository, err = repository.NewRepository(repositoryConfig)
	d.addCloser(d.repository.Close)
	return err
}

// initAuthClient ...
func (d *deps) initAuthClient(ctx context.Context) (err error) {
	keyfilePath := d.config.FirebaseConfig.FirebaseKeyFilePath
	d.authClient, err = authClient.NewAuthClient(keyfilePath)
	return err
}

// initPayClient ...
func (d *deps) initPayClient(ctx context.Context) (err error) {
	d.payClient, err = payClient.NewPayClient(d.logger, d.config.PaymentURLExpirationPeriod)
	return err
}

// initRabbitMq ...
func (d *deps) initRabbitMq(ctx context.Context) (err error) {
	// config
	logger := d.logger
	rabbitMQConfig := &d.config.RabbitMQConfig
	config := rabbitmq.RabbitMqClientConfig{
		Logger:         logger,
		RabbitMQConfig: rabbitMQConfig,
	}

	// init
	d.rabbitMqClient, err = rabbitmq.NewRabbitMqClient(config)
	if err != nil {
		d.logger.Fatalln("new rabbit conn: ", err)
	}
	d.logger.Debug("connected to rabbit")

	return nil
}

// initLogic ...
func (d *deps) initLeaWashPublisher(ctx context.Context) (err error) {
	d.leaWashPublisher, err = leawash.NewLeaWashPublisher(d.rabbitMqClient)
	if err != nil {
		return err
	}

	return nil
}

// initLogic ...
func (d *deps) initLogic(ctx context.Context) (err error) {
	conf := logic.LogicConfig{
		Logger:                       d.logger,
		NotificationExpirationPeriod: d.config.NotificationExpirationPeriod,
		Repository:                   d.repository,
		LeaWashPublisher:             d.leaWashPublisher,
		PayClient:                    d.payClient,
		AuthClient:                   d.authClient,
		BrokerUserCreator:            d.brokerUserCreator,
		PasswordLength:               d.config.PasswordLength,
	}

	d.logic, err = logic.NewLogic(ctx, conf)
	if err != nil {
		return err
	}

	return err
}

// initLeaWashConsumer ...
func (d *deps) initLeaWashConsumer(ctx context.Context) (err error) {
	d.leaWashConsumer, err = leawash.NewLeaWashConsumer(d.logger, d.rabbitMqClient, d.logic, d.leaWashPublisher)
	if err != nil {
		d.logger.Fatalln("new rabbit conn: ", err)
	}
	d.logger.Debug("connected to rabbit")

	return nil
}

// initBrokerUserCreator ...
func (d *deps) initBrokerUserCreator(ctx context.Context) (err error) {
	d.brokerUserCreator, err = leawash.NewBrokerUserCreator(d.rabbitMqClient)
	if err != nil {
		d.logger.Fatalln("new rabbit conn: ", err)
	}
	d.logger.Debug("connected to rabbit")

	return nil
}

// initHttpApi ...
func (d *deps) initHttpApi(ctx context.Context) (err error) {
	restApiConfig := restApi.RestApiConfig{
		Logger: d.logger,
		Logic:  d.logic,
	}
	d.api, err = restApi.NewApi(restApiConfig)
	return err
}

// initHttpServer ...
func (d *deps) initHttpServer(ctx context.Context) (err error) {
	serverConfig := server.ServerConfig{
		Logger:         d.logger,
		Host:           d.config.Host,
		Port:           d.config.HTTPPort,
		AllowedOrigins: d.config.AllowedOrigins,
		Api:            d.api,
	}
	d.httpServer, err = server.NewServer(serverConfig)
	return err
}

// initTimerTriggeredScheduler ...
func (d *deps) initTimerTriggeredScheduler(ctx context.Context) (err error) {
	paymentTiker := time.NewTicker(d.config.PaymentSyncTimeOut)
	d.timerTriggeredScheduler, err = timerTriggeredScheduler.NewTimerTriggeredScheduler(d.logger, d.logic, paymentTiker)
	return err
}
