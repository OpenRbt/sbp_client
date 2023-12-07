package config

import (
	"errors"
	"sbp/internal/app"
	"sbp/internal/infrastructure/firebase"
	"time"

	"go.uber.org/zap"
)

type (
	Config struct {
		HTTPPort                     string        `env:"HTTP_PORT" envDefault:"8080"`
		BasePath                     string        `env:"HTTP_BASE_PATH" envDefault:""`
		AllowedOrigins               string        `env:"HTTP_ALLOWED_ORIGINS" envDefault:"*"`
		Host                         string        `env:"HTTP_HOST"  envDefault:""`
		LogLevel                     string        `env:"LOG_LEVEL" envDefault:""`
		PasswordLength               int           `env:"PASSWORD_LENGTH" envDefault:"10"`
		PaymentSyncTimeOut           time.Duration `env:"PAYMENT_SYNC_TIMEOUT" envDefault:"5m"`
		NotificationExpirationPeriod time.Duration `env:"NOTIFICATION_EXPIRATION_PERIOD" envDefault:"1h"`
		PaymentURLExpirationPeriod   time.Duration `env:"PAYMENT_URL_EXPIRATION_PERIOD" envDefault:"1h"`
		DB                           DBConfig
		FirebaseConfig               FirebaseConfig
		RabbitMQConfig               RabbitMQConfig
	}

	DBConfig struct {
		Host           string `env:"DB_HOST" envDefault:"localhost"`
		Port           string `env:"DB_PORT" envDefault:"8092"`
		Database       string `env:"DB_DATABASE" envDefault:"wash_sbp_db"`
		User           string `env:"DB_USER" envDefault:"wash_admin"`
		Password       string `env:"DB_PASSWORD" envDefault:"wash_admin"`
		MigrationsPath string `env:"MIGRATIONS_PATH" envDefault:"internal/repository/migrations"`
	}

	FirebaseConfig struct {
		FirebaseKeyFilePath string `env:"FB_KEYFILE_PATH" envDefault:"/home/roman/Документы/sbp/sbp/environment/firebase/fb_key.json"`
	}

	RabbitMQConfig struct {
		Url      string `env:"RABBIT_SERVICE_URL" envDefault:"localhost"`
		Port     string `env:"RABBIT_SERVICE_PORT" envDefault:"5672"`
		PortWeb  string `env:"RABBIT_SERVICE_PORT_WEB" envDefault:"15672"`
		User     string `env:"RABBIT_SERVICE_USER" envDefault:"sbp_admin"`
		Password string `env:"RABBIT_SERVICE_PASSWORD" envDefault:"sbp_admin"`
		Secure   bool   `env:"RABBIT_SERVICE_SECURE" envDefault:"true"`
	}

	RabbitMqClientConfig struct {
		Logger         *zap.SugaredLogger
		RabbitMQConfig *RabbitMQConfig
	}

	ServiceConfig struct {
		Logger                       *zap.SugaredLogger
		NotificationExpirationPeriod time.Duration
		PasswordLength               int
		Repository                   app.Repository
		LeaWashPublisher             app.LeaWashPublisher
		PayClient                    app.PaymentClient
		BrokerUserCreator            app.UserBroker
	}

	RestApiConfig struct {
		Logger   *zap.SugaredLogger
		Svc      app.Service
		Firebase *firebase.FirebaseClient
	}

	RepositoryConfig struct {
		DBConfig *DBConfig
		Logger   *zap.SugaredLogger
	}

	ServerConfig struct {
		Logger         *zap.SugaredLogger
		Host           string
		Port           string
		AllowedOrigins string
		Api            app.Api
	}
)

func (conf *RabbitMqClientConfig) CheckRabbitMqClientConfig() error {
	if conf.Logger == nil {
		return errors.New("logger is empty")
	}
	if conf.RabbitMQConfig == nil {
		return errors.New("rabbit_mq_config is empty")
	}
	return nil
}

func (conf *RepositoryConfig) CheckRepositoryConfig() error {
	if conf.DBConfig == nil {
		return errors.New("repository db config is empty")
	}
	if conf.Logger == nil {
		return errors.New("repository logger is empty")
	}

	return nil
}

func (conf *ServiceConfig) CheckServiceConfig() error {
	if conf.Logger == nil {
		return errors.New("service logger is empty")
	}
	if conf.Repository == nil {
		return errors.New("service repository is empty")
	}
	if conf.LeaWashPublisher == nil {
		return errors.New("service lea_wash_publisher is empty")
	}
	if conf.PayClient == nil {
		return errors.New("service pay_client is empty")
	}
	if conf.NotificationExpirationPeriod == 0 {
		return errors.New("service notification_expiration_period is 0")
	}
	if conf.BrokerUserCreator == nil {
		return errors.New("service broker_user_creator is nil")
	}
	return nil
}

func (conf *RestApiConfig) CheckRestApiConfig() error {
	if conf.Logger == nil {
		return errors.New("api logger is empty")
	}
	if conf.Svc == nil {
		return errors.New("api logic is empty")
	}

	return nil
}

func (conf *ServerConfig) CheckServerConfig() error {
	if conf.Host == "" {
		return errors.New("server host is empty")
	}
	if conf.Port == "" {
		return errors.New("server port is empty")
	}
	if conf.AllowedOrigins == "" {
		return errors.New("server allowedOrigins is empty")
	}
	if conf.Api == nil {
		return errors.New("server api is empty")
	}
	return nil
}
