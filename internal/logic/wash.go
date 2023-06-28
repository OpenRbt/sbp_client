package logic

import (
	"context"
	logicEntities "sbp/internal/logic/entities"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// WashServerLogic ...
type WashServerLogic struct {
	logger     *zap.SugaredLogger
	repository WashServerRepository
}

// WashServerRepository ...
type WashServerRepository interface {
	CreateWashServer(ctx context.Context, admin uuid.UUID, newServer logicEntities.RegisterWashServer) (logicEntities.WashServer, error)
	UpdateWashServer(ctx context.Context, updateWashServer logicEntities.UpdateWashServer) error
	DeleteWashServer(ctx context.Context, id uuid.UUID) error
	GetWashServer(ctx context.Context, id uuid.UUID) (logicEntities.WashServer, error)
	GetWashServerList(ctx context.Context, pagination logicEntities.Pagination) ([]logicEntities.WashServer, error)
}

// newWashServerLogic ...
func newWashServerLogic(ctx context.Context, logger *zap.SugaredLogger, repository WashServerRepository) (*WashServerLogic, error) {
	logic := WashServerLogic{
		logger:     logger,
		repository: repository,
	}

	return &logic, nil
}

// CreateWashServer ...
func (logic *WashServerLogic) CreateWashServer(ctx context.Context, admin logicEntities.SbpAdmin, newServer logicEntities.RegisterWashServer) (logicEntities.WashServer, error) {
	w, err := logic.repository.CreateWashServer(ctx, admin.ID, newServer)
	if err != nil {
		return logicEntities.WashServer{}, err
	}

	return w, nil
}

// GetWashServer ...
func (logic *WashServerLogic) GetWashServer(ctx context.Context, id uuid.UUID) (logicEntities.WashServer, error) {
	washServer, err := logic.repository.GetWashServer(ctx, id)
	if err != nil {
		return logicEntities.WashServer{}, err
	}

	return washServer, nil
}

// UpdateWashServer ...
func (logic *WashServerLogic) UpdateWashServer(ctx context.Context, admin logicEntities.SbpAdmin, updateWashServer logicEntities.UpdateWashServer) error {
	washServer, err := logic.repository.GetWashServer(ctx, updateWashServer.ID)
	if err != nil {
		return err
	}
	if washServer.Owner != admin.ID {
		return logicEntities.ErrUserNotOwner
	}
	err = logic.repository.UpdateWashServer(ctx, updateWashServer)
	if err != nil {
		return err
	}
	return nil
}

// DeleteWashServer ...
func (logic *WashServerLogic) DeleteWashServer(ctx context.Context, admin logicEntities.SbpAdmin, id uuid.UUID) error {
	washServer, err := logic.repository.GetWashServer(ctx, id)
	if err != nil {
		return err
	}
	if washServer.Owner != admin.ID {
		return logicEntities.ErrUserNotOwner
	}

	err = logic.repository.DeleteWashServer(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetWashServerList ...
func (logic *WashServerLogic) GetWashServerList(ctx context.Context, pagination logicEntities.Pagination) ([]logicEntities.WashServer, error) {
	washServers, err := logic.repository.GetWashServerList(ctx, pagination)
	if err != nil {
		return nil, err
	}

	return washServers, nil
}
