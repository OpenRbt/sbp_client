package logic

import (
	"context"
	"errors"
	"sbp/pkg/bootstrap"
	"time"

	"go.uber.org/zap"
)

// limitToReadWashes ...
const (
	layer = "logic"
)

// Repository ...
type Repository interface {
	// Pay
	PayRepository
	// User ...
	UserRepository
	// Wash ...
	WashRepository
	// Close ...
	Close() error
}

// Logic ...
type Logic struct {
	logger *zap.SugaredLogger
	AuthLogic
	UserLogic
	WashLogic
	PaymentLogic
}

// LogicConfig ...
type LogicConfig struct {
	Logger                       *zap.SugaredLogger
	NotificationExpirationPeriod time.Duration
	PasswordLength               int
	Repository                   Repository
	LeaWashPublisher             LeaWashPublisher
	PayClient                    PayClient
	AuthClient                   AuthClient
	BrokerUserCreator            BrokerUserCreator
}

// CheckLogicConfig ...
func CheckLogicConfig(conf LogicConfig) error {
	if conf.Logger == nil {
		return errors.New("logic logger is empty")
	}
	if conf.Repository == nil {
		return errors.New("logic repository is empty")
	}
	if conf.LeaWashPublisher == nil {
		return errors.New("logic lea_wash_publisher is empty")
	}
	if conf.PayClient == nil {
		return errors.New("logic pay_client is empty")
	}
	if conf.AuthClient == nil {
		return errors.New("logic auth_client is empty")
	}
	if conf.NotificationExpirationPeriod == 0 {
		return errors.New("logic notification_expiration_period is 0")
	}
	if conf.BrokerUserCreator == nil {
		return errors.New("logic broker_user_creator is nil")
	}
	return nil
}

// NewLogic ...
func NewLogic(ctx context.Context, config LogicConfig) (*Logic, error) {
	err := CheckLogicConfig(config)
	if err != nil {
		return nil, bootstrap.CustomError(layer, "CheckLogicConfig", err)
	}

	authLogic, err := newAuthLogic(ctx, config.Logger, config.AuthClient)
	if err != nil {
		return nil, err
	}

	userLogic, err := newUserLogic(ctx, config.Logger, config.Repository)
	if err != nil {
		return nil, err
	}

	washLogic, err := newWashLogic(ctx, config.Logger, config.Repository, config.BrokerUserCreator, config.PasswordLength)
	if err != nil {
		return nil, err
	}

	paymentLogic, err := newPaymentLogic(
		ctx,
		config.Logger,
		config.NotificationExpirationPeriod,
		config.PayClient,
		config.Repository,
		config.LeaWashPublisher,
		washLogic,
	)
	if err != nil {
		return nil, err
	}

	logic := Logic{
		logger:       config.Logger,
		AuthLogic:    *authLogic,
		UserLogic:    *userLogic,
		WashLogic:    *washLogic,
		PaymentLogic: *paymentLogic,
	}

	return &logic, nil
}
