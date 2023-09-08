package repository

import (
	"context"
	"errors"
	logicEntities "sbp/internal/logic/entities"
	repConverter "sbp/internal/repository/converter"
	repEntities "sbp/internal/repository/entities"

	"github.com/gocraft/dbr/v2"
)

// GetUser ...
func (s *Repository) GetUser(ctx context.Context, identity string) (logicEntities.User, error) {
	var dbUser repEntities.User

	err := s.db.
		NewSession(nil).
		Select("*").
		From("users").
		Where("identity_uid = ?", identity).
		LoadOneContext(ctx, &dbUser)

	switch {
	case err == nil:
		return repConverter.ConvertUserFromDB(dbUser), err
	case errors.Is(err, dbr.ErrNotFound):
		return logicEntities.User{}, logicEntities.ErrNotFound
	default:
		return logicEntities.User{}, err
	}
}

// CreateUser ...
func (s *Repository) CreateUser(ctx context.Context, identity string) (logicEntities.User, error) {
	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return logicEntities.User{}, err
	}

	var dbUser repEntities.User
	err = tx.
		InsertInto("users").
		Columns("identity_uid").
		Values(identity).
		Returning("id", "identity_uid").
		LoadContext(ctx, &dbUser)

	if err != nil {
		return logicEntities.User{}, err
	}

	return repConverter.ConvertUserFromDB(dbUser), tx.Commit()
}

// GetOrCreateAdminIfNotExists ...
func (s *Repository) GetOrCreateAdminIfNotExists(ctx context.Context, identity string) (logicEntities.User, error) {
	dbUser, err := s.GetUser(ctx, identity)

	if err != nil {
		if errors.Is(err, logicEntities.ErrNotFound) {
			return s.CreateUser(ctx, identity)
		}

		return logicEntities.User{}, err
	}

	return dbUser, err
}
