package main

import (
	"context"
	"fmt"
	"time"

	restApi "sbp/internal/api"
	"sbp/internal/app"
	"sbp/internal/config"
	"sbp/internal/infrastructure/firebase"
	payClient "sbp/internal/infrastructure/pay-client"
	leawash "sbp/internal/infrastructure/rabbit/lea"
	shareRabbit "sbp/internal/infrastructure/rabbit/share"
	repository "sbp/internal/repository"
	"sbp/internal/services"
	timetriggeredtasks "sbp/internal/time-triggered-scheduler"
	"sbp/pkg/bootstrap"
	"sbp/pkg/rabbitmq"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type HttpServer interface {
	Run() error
}

type deps struct {
	config *config.Config
	logger *zap.SugaredLogger

	repository       app.Repository
	rabbitMqClient   *rabbitmq.RabbitMqClient
	payClient        app.PaymentClient
	leaWashPublisher app.LeaWashPublisher
	leaWashConsumer  Consumer
	sharePublisher   app.SharePublisher
	shareConsumer    Consumer
	firebase         *firebase.FirebaseClient

	svc app.Service

	api        app.Api
	httpServer HttpServer

	timeTriggeredScheduler timetriggeredtasks.TimeTriggeredScheduler

	closeFuncs []func() error
}

type Consumer interface {
	Close()
}

func InitDeps(ctx context.Context, config *config.Config, logger *zap.SugaredLogger) (*deps, error) {
	d := deps{
		config: config,
		logger: logger,
	}

	inits := []func(context.Context) error{
		d.initRabbitMq,
		d.initRepository,
		d.initFirebase,
		d.initPayClient,
		d.initLeaWashPublisher,
		d.initSharePublisher,
		d.initServices,
		d.initShareConsumer,
		d.initLeaWashConsumer,
		d.initHttpApi,
		d.initHttpServer,
		d.initTimeTriggeredScheduler,
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

func (d *deps) close() {
	closeFuncs := []func(){
		d.leaWashConsumer.Close,
	}
	d.logger.Info("start close deps")
	for _, f := range closeFuncs {
		f()
	}
}

func (d *deps) addCloser(f func() error) {
	d.closeFuncs = append(d.closeFuncs, f)
}

func (d *deps) initRepository(ctx context.Context) (err error) {
	repositoryConfig := config.RepositoryConfig{
		DBConfig: &d.config.DB,
		Logger:   d.logger,
	}
	d.repository, err = repository.NewRepository(repositoryConfig)
	d.addCloser(d.repository.Close)
	return err
}

func (d *deps) initFirebase(ctx context.Context) (err error) {
	d.firebase, err = firebase.NewAuthClient(d.config.FirebaseConfig.FirebaseKeyFilePath, d.repository)
	return err
}

func (d *deps) initPayClient(ctx context.Context) (err error) {
	d.payClient, err = payClient.NewPayClient(d.logger, d.config.PaymentURLExpirationPeriod)
	return err
}

func (d *deps) initRabbitMq(ctx context.Context) (err error) {
	logger := d.logger
	rabbitMQConfig := &d.config.RabbitMQConfig
	config := config.RabbitMqClientConfig{
		Logger:         logger,
		RabbitMQConfig: rabbitMQConfig,
	}

	d.rabbitMqClient, err = rabbitmq.NewRabbitMqClient(config)
	if err != nil {
		d.logger.Fatalln("new rabbit conn: ", err)
	}
	d.logger.Debug("connected to rabbit")

	return nil
}

func (d *deps) initServices(ctx context.Context) (err error) {
	conf := config.ServiceConfig{
		Logger:                       d.logger,
		NotificationExpirationPeriod: d.config.NotificationExpirationPeriod,
		Repository:                   d.repository,
		LeaWashPublisher:             d.leaWashPublisher,
		PayClient:                    d.payClient,
		SharePublisher:               d.sharePublisher,
		PasswordLength:               d.config.PasswordLength,
	}

	d.svc, err = services.NewServices(ctx, conf)
	if err != nil {
		return err
	}

	return err
}

func (d *deps) initLeaWashPublisher(ctx context.Context) (err error) {
	d.leaWashPublisher, err = leawash.NewLeaWashPublisher(d.logger, d.rabbitMqClient)
	if err != nil {
		return err
	}

	return nil
}

func (d *deps) initLeaWashConsumer(ctx context.Context) (err error) {
	d.leaWashConsumer, err = leawash.NewLeaConsumer(d.logger, d.rabbitMqClient, d.svc, d.leaWashPublisher)
	if err != nil {
		d.logger.Fatalln("new rabbit conn: ", err)
	}
	d.logger.Debug("connected to rabbit")

	return nil
}

func (d *deps) initSharePublisher(ctx context.Context) (err error) {
	d.sharePublisher, err = shareRabbit.NewSharePublisher(d.logger, d.rabbitMqClient)
	if err != nil {
		return err
	}

	err = d.sharePublisher.SendDataRequest()
	if err != nil {
		d.logger.Fatalln("unable to send data request: ", err)
	}

	return nil
}

func (d *deps) initShareConsumer(ctx context.Context) (err error) {
	d.leaWashConsumer, err = shareRabbit.NewShareConsumer(d.logger, d.rabbitMqClient, d.svc, d.sharePublisher)
	if err != nil {
		d.logger.Fatalln("new rabbit conn: ", err)
	}
	d.logger.Debug("connected to rabbit")

	return nil
}

func (d *deps) initHttpApi(ctx context.Context) (err error) {
	restApiConfig := config.RestApiConfig{
		Logger:   d.logger,
		Svc:      d.svc,
		Firebase: d.firebase,
	}
	d.api, err = restApi.NewApi(restApiConfig)
	return err
}

func (d *deps) initHttpServer(ctx context.Context) (err error) {
	serverConfig := config.ServerConfig{
		Logger:         d.logger,
		Host:           d.config.Host,
		Port:           d.config.HTTPPort,
		AllowedOrigins: d.config.AllowedOrigins,
		Api:            d.api,
	}
	d.httpServer, err = restApi.NewServer(serverConfig)
	return err
}

func (d *deps) initTimeTriggeredScheduler(ctx context.Context) (err error) {
	paymentTiker := time.NewTicker(d.config.PaymentSyncTimeOut)
	d.timeTriggeredScheduler, err = timetriggeredtasks.NewTimeTriggeredScheduler(d.logger, d.svc, paymentTiker)
	return err
}

type App struct {
	deps *deps
}

func NewApp(ctx context.Context, envFilePath string) (*App, error) {
	config, err := getConfig(envFilePath)
	if err != nil {
		return nil, err
	}

	logger, err := getLogger(config.LogLevel)
	if err != nil {
		return nil, err
	}

	deps, err := InitDeps(ctx, config, logger)
	if err != nil {
		return nil, err
	}

	return &App{
		deps: deps,
	}, nil
}

func (a App) Run() error {
	return a.deps.httpServer.Run()
}

func (a App) Close() {
	a.deps.close()
}

func getConfig(envFilePath string) (*config.Config, error) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %s", err.Error())
	}

	config, err := bootstrap.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("new config: %s", err.Error())
	}

	return config, nil
}

func getLogger(logLevel string) (l *zap.SugaredLogger, err error) {
	logger, err := bootstrap.NewLogger(logLevel)
	if err != nil {
		return nil, fmt.Errorf("new logger: %s", err.Error())
	}
	return logger, nil
}
