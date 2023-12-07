package app

import (
	"sbp/internal/entities"

	uuid "github.com/satori/go.uuid"
)

type (
	OrganizationRepository interface {
		GetOrganizationByID(ctx Ctx, id uuid.UUID) (entities.Organization, error)
		InsertOrganization(ctx Ctx, org entities.Organization) error
	}

	OrganizationService interface {
		UpsertOrganization(ctx Ctx, org entities.Organization) error
	}
)
