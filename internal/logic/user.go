package logic

import (
	"context"
	"errors"
	logicEntities "sbp/internal/logic/entities"

	"go.uber.org/zap"
)

// UserLogic ...
type UserLogic struct {
	logger     *zap.SugaredLogger
	repository UserRepository
}

// UserRepository ...
type UserRepository interface {
	CreateUser(ctx context.Context, identity string) (logicEntities.User, error)
	GetUser(ctx context.Context, identity string) (logicEntities.User, error)
}

// newUserLogic ...
func newUserLogic(ctx context.Context, logger *zap.SugaredLogger, repository UserRepository) (*UserLogic, error) {
	return &UserLogic{
		logger:     logger,
		repository: repository,
	}, nil
}

// CreateUser ...
func (logic *UserLogic) CreateUser(ctx context.Context, auth *logicEntities.Auth) (*logicEntities.User, error) {
	admin, err := logic.repository.CreateUser(ctx, auth.UID)
	if err != nil {
		return nil, err
	}

	if admin.ID.String() == "0" {
		return nil, errors.New("admin.ID = 0")
	}

	return &admin, nil
}

// GetUser ...
func (logic *UserLogic) GetUser(ctx context.Context, auth *logicEntities.Auth) (*logicEntities.User, error) {
	admin, err := logic.repository.GetUser(ctx, auth.UID)
	if err != nil {
		return nil, err
	}

	if admin.ID.String() == "0" {
		return nil, errors.New("admin.ID = 0")
	}

	return &admin, nil
}
