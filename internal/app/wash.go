package app

import (
	"sbp/internal/entities"

	uuid "github.com/satori/go.uuid"
)

type (
	WashRepository interface {
		CreateWash(ctx Ctx, new entities.WashCreation) (entities.Wash, error)
		UpdateWash(ctx Ctx, id uuid.UUID, updateWash entities.WashUpdate) error
		DeleteWash(ctx Ctx, id uuid.UUID) error
		GetWashByID(ctx Ctx, id uuid.UUID) (entities.Wash, error)
		GetWashes(ctx Ctx, filter entities.WashFilter) ([]entities.Wash, error)

		AssignWashToGroup(ctx Ctx, washID, groupID uuid.UUID) error
	}

	WashService interface {
		CreateWash(ctx Ctx, auth *Auth, newWash entities.WashCreation) (entities.Wash, error)
		UpdateWash(ctx Ctx, auth *Auth, id uuid.UUID, updateWash entities.WashUpdate) error
		DeleteWash(ctx Ctx, auth *Auth, id uuid.UUID) error
		GetWashByID(ctx Ctx, auth *Auth, id uuid.UUID) (entities.Wash, error)
		GetWashes(ctx Ctx, auth *Auth, filter entities.WashFilter) ([]entities.Wash, error)

		AssignWashToGroup(ctx Ctx, auth *Auth, washID, groupID uuid.UUID) error
	}
)
