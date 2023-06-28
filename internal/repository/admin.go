package repository

import (
	"context"
	"errors"
	logicEntities "sbp/internal/logic/entities"
	repConverter "sbp/internal/repository/converter"
	repEntities "sbp/internal/repository/entities"

	"github.com/gocraft/dbr/v2"
)

// GetSbpAdmin ...
func (s *Repository) GetSbpAdmin(ctx context.Context, identity string) (logicEntities.SbpAdmin, error) {
	var dbSbpAdmin repEntities.SbpAdmin

	err := s.db.
		NewSession(nil).
		Select("*").
		From("users").
		Where("identity_uid = ?", identity).
		LoadOneContext(ctx, &dbSbpAdmin)

	switch {
	case err == nil:
		return repConverter.ConvertSbpAdminFromDB(dbSbpAdmin), err
	case errors.Is(err, dbr.ErrNotFound):
		return logicEntities.SbpAdmin{}, logicEntities.ErrNotFound
	default:
		return logicEntities.SbpAdmin{}, err
	}
}

// CreateSbpAdmin ...
func (s *Repository) CreateSbpAdmin(ctx context.Context, identity string) (logicEntities.SbpAdmin, error) {
	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return logicEntities.SbpAdmin{}, err
	}

	var dbSbpAdmin repEntities.SbpAdmin
	err = tx.
		InsertInto("users").
		Columns("identity_uid").
		Values(identity).
		Returning("id", "identity_uid").
		LoadContext(ctx, &dbSbpAdmin)

	if err != nil {
		return logicEntities.SbpAdmin{}, err
	}

	return repConverter.ConvertSbpAdminFromDB(dbSbpAdmin), tx.Commit()
}

// GetOrCreateAdminIfNotExists ...
func (s *Repository) GetOrCreateAdminIfNotExists(ctx context.Context, identity string) (logicEntities.SbpAdmin, error) {
	dbSbpAdmin, err := s.GetSbpAdmin(ctx, identity)

	if err != nil {
		if errors.Is(err, logicEntities.ErrNotFound) {
			return s.CreateSbpAdmin(ctx, identity)
		}

		return logicEntities.SbpAdmin{}, err
	}

	return dbSbpAdmin, err
}
