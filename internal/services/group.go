package services

import (
	"context"
	"errors"
	"sbp/internal/app"
	"sbp/internal/entities"

	"go.uber.org/zap"
)

type groupService struct {
	logger     *zap.SugaredLogger
	repository app.GroupRepository
}

func NewGroupService(ctx context.Context, logger *zap.SugaredLogger, repository app.GroupRepository) (*groupService, error) {
	return &groupService{
		logger:     logger,
		repository: repository,
	}, nil
}

func (svc *groupService) UpsertGroup(ctx context.Context, newGroup entities.Group) error {
	curGroup, err := svc.repository.GetGroupByID(ctx, newGroup.ID)
	if err != nil && !errors.Is(err, entities.ErrNotFound) {
		return err
	}

	if !errors.Is(err, entities.ErrNotFound) && newGroup.Version <= curGroup.Version {
		return entities.ErrBadVersion
	}

	err = svc.repository.InsertGroup(ctx, newGroup)
	if err != nil {
		return err
	}

	return nil
}
