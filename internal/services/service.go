package services

import (
	"context"
	"sbp/internal/config"
	"sbp/internal/helpers"

	"go.uber.org/zap"
)

type Services struct {
	logger *zap.SugaredLogger
	userService
	washService
	paymentService
	organizationService
	groupService
}

func NewServices(ctx context.Context, cfg config.ServiceConfig) (*Services, error) {
	err := cfg.CheckServiceConfig()
	if err != nil {
		return nil, helpers.CustomError("app", "CheckLogicConfig", err)
	}

	userService, err := NewUserService(ctx, cfg.Logger, cfg.Repository)
	if err != nil {
		return nil, err
	}

	washService, err := NewWashService(ctx, cfg.Logger, cfg.Repository, cfg.BrokerUserCreator, cfg.PasswordLength)
	if err != nil {
		return nil, err
	}

	orgService, err := NewOrganizationService(ctx, cfg.Logger, cfg.Repository)
	if err != nil {
		return nil, err
	}

	groupService, err := NewGroupService(ctx, cfg.Logger, cfg.Repository)
	if err != nil {
		return nil, err
	}

	paymentService, err := NewPaymentService(
		ctx,
		cfg.Logger,
		cfg.NotificationExpirationPeriod,
		cfg.PayClient,
		cfg.Repository,
		cfg.LeaWashPublisher,
		washService,
	)
	if err != nil {
		return nil, err
	}

	services := Services{
		logger:              cfg.Logger,
		userService:         *userService,
		washService:         *washService,
		paymentService:      *paymentService,
		organizationService: *orgService,
		groupService:        *groupService,
	}

	return &services, nil
}
