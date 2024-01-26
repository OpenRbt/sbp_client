package services

import (
	"context"
	"errors"
	"sbp/internal/app"
	"sbp/internal/entities"

	"go.uber.org/zap"
)

type userService struct {
	logger     *zap.SugaredLogger
	repository app.UserRepository
}

func NewUserService(ctx context.Context, logger *zap.SugaredLogger, repository app.UserRepository) (*userService, error) {
	return &userService{
		logger:     logger,
		repository: repository,
	}, nil
}

func (svc *userService) UpsertUser(ctx context.Context, newUser entities.User) error {
	curUser, err := svc.repository.GetUserByID(ctx, newUser.ID)
	if err != nil && !errors.Is(err, entities.ErrNotFound) {
		return err
	}

	if !errors.Is(err, entities.ErrNotFound) && newUser.Version <= curUser.Version {
		return entities.ErrBadVersion
	}

	err = svc.repository.InsertUser(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}
