package bootstrap

import (
	"errors"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
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
	// Sync ...
}

type DBConfig struct {
	Host           string `env:"DB_HOST" envDefault:"localhost"`
	Port           string `env:"DB_PORT" envDefault:"8092"`
	Database       string `env:"DB_DATABASE" envDefault:"sbp_db"`
	User           string `env:"DB_USER" envDefault:"wash_admin"`
	Password       string `env:"DB_PASSWORD" envDefault:"wash_admin"`
	MigrationsPath string `env:"MIGRATIONS_PATH" envDefault:"internal/repository/migrations"`
}

type FirebaseConfig struct {
	FirebaseKeyFilePath string `env:"FB_KEYFILE_PATH" envDefault:"/home/roman/Документы/sbp/sbp/environment/firebase/fb_key.json"`
}

type RabbitMQConfig struct {
	Url      string `env:"RABBIT_SERVICE_URL" envDefault:"localhost"`
	Port     string `env:"RABBIT_SERVICE_PORT" envDefault:"5672"`
	PortWeb  string `env:"RABBIT_SERVICE_PORT_WEB" envDefault:"15672"`
	User     string `env:"RABBIT_SERVICE_USER" envDefault:"sbp_admin"`
	Password string `env:"RABBIT_SERVICE_PASSWORD" envDefault:"sbp_admin"`
	Secure   bool   `env:"RABBIT_SERVICE_SECURE" envDefault:"true"`
}

func NewConfig(configFiles ...string) (*Config, error) {
	var c Config
	err := godotenv.Load(configFiles...)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	return &c, env.Parse(&c, env.Options{
		RequiredIfNoDef: true,
	})

}
