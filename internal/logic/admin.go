package logic

import (
	"context"
	"errors"
	logicEntities "sbp/internal/logic/entities"

	"go.uber.org/zap"
)

// AdminLogic ...
type AdminLogic struct {
	logger     *zap.SugaredLogger
	repository AdminRepository
}

// AdminRepository ...
type AdminRepository interface {
	GetOrCreateAdminIfNotExists(ctx context.Context, identity string) (logicEntities.SbpAdmin, error)
	GetSbpAdmin(ctx context.Context, identity string) (logicEntities.SbpAdmin, error)
}

// newAdminLogic ...
func newAdminLogic(ctx context.Context, logger *zap.SugaredLogger, repository AdminRepository) (*AdminLogic, error) {
	return &AdminLogic{
		logger:     logger,
		repository: repository,
	}, nil
}

// GetOrCreateAdminIfNotExists ...
func (logic *AdminLogic) GetOrCreateAdminIfNotExists(ctx context.Context, auth *logicEntities.Auth) (*logicEntities.SbpAdmin, error) {
	admin, err := logic.repository.GetOrCreateAdminIfNotExists(ctx, auth.UID)
	if err != nil {
		return nil, err
	}

	if admin.ID.String() == "0" {
		return nil, errors.New("admin.ID = 0")
	}

	return &admin, nil
}

// GetSbpAdmin ...
func (logic *AdminLogic) GetSbpAdmin(ctx context.Context, auth *logicEntities.Auth) (*logicEntities.SbpAdmin, error) {
	admin, err := logic.repository.GetSbpAdmin(ctx, auth.UID)
	if err != nil {
		return nil, err
	}

	if admin.ID.String() == "0" {
		return nil, errors.New("admin.ID = 0")
	}

	return &admin, nil
}
