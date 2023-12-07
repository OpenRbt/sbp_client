package app

import (
	"sbp/internal/entities"
)

type (
	UserRepository interface {
		GetUserByID(ctx Ctx, id string) (entities.User, error)
		InsertUser(ctx Ctx, user entities.User) error
	}

	UserService interface {
		UpsertUser(ctx Ctx, user entities.User) error
	}

	UserBroker interface {
		CreateUser(login string, password string) error
	}
)
