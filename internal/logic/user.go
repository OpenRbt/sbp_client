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
	GetOrCreateUser(ctx context.Context, identity string) (logicEntities.User, error)
}

// newUserLogic ...
func newUserLogic(ctx context.Context, logger *zap.SugaredLogger, repository UserRepository) (*UserLogic, error) {
	return &UserLogic{
		logger:     logger,
		repository: repository,
	}, nil
}

// GetOrCreateUser ...
func (logic *UserLogic) GetOrCreateUser(ctx context.Context, auth *logicEntities.Auth) (*logicEntities.User, error) {
	admin, err := logic.repository.GetOrCreateUser(ctx, auth.UID)
	if err != nil {
		return nil, err
	}

	if admin.ID.String() == "0" {
		return nil, errors.New("admin.ID = 0")
	}

	return &admin, nil
}
