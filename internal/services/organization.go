package services

import (
	"context"
	"errors"
	"sbp/internal/app"
	"sbp/internal/entities"

	"go.uber.org/zap"
)

type organizationService struct {
	logger     *zap.SugaredLogger
	repository app.OrganizationRepository
}

func NewOrganizationService(ctx context.Context, logger *zap.SugaredLogger, repository app.OrganizationRepository) (*organizationService, error) {
	return &organizationService{
		logger:     logger,
		repository: repository,
	}, nil
}

func (svc *organizationService) UpsertOrganization(ctx context.Context, newOrg entities.Organization) error {
	curOrg, err := svc.repository.GetOrganizationByID(ctx, newOrg.ID)
	if err != nil && !errors.Is(err, entities.ErrNotFound) {
		return err
	}

	if !errors.Is(err, entities.ErrNotFound) && newOrg.Version <= curOrg.Version {
		return entities.ErrBadVersion
	}

	err = svc.repository.InsertOrganization(ctx, newOrg)
	if err != nil {
		return err
	}

	return nil
}
