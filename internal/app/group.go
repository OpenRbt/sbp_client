package app

import (
	"sbp/internal/entities"

	uuid "github.com/satori/go.uuid"
)

type (
	GroupRepository interface {
		GetGroupByID(ctx Ctx, id uuid.UUID) (entities.Group, error)
		InsertGroup(ctx Ctx, group entities.Group) error
	}

	GroupService interface {
		UpsertGroup(ctx Ctx, group entities.Group) error
	}
)
