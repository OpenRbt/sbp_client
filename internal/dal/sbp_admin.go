package dal

import (
	"context"
	"errors"
	"sbp/internal/app"
	"sbp/internal/conversions"
	"sbp/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
)

func (s *Storage) GetSbpAdmin(ctx context.Context, identity string) (app.SbpAdmin, error) {
	var dbSbpAdmin dbmodels.SbpAdmin

	err := s.db.NewSession(nil).
		Select("*").
		From("users").
		Where("identity_uid = ?", identity).
		LoadOneContext(ctx, &dbSbpAdmin)

	switch {
	case err == nil:
		return conversions.SbpAdminFromDB(dbSbpAdmin), err
	case errors.Is(err, dbr.ErrNotFound):
		return app.SbpAdmin{}, app.ErrNotFound
	default:
		return app.SbpAdmin{}, err
	}
}

func (s *Storage) CreateSbpAdmin(ctx context.Context, identity string) (app.SbpAdmin, error) {
	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return app.SbpAdmin{}, err
	}

	var dbSbpAdmin dbmodels.SbpAdmin
	err = tx.
		InsertInto("users").
		Columns("identity_uid").
		Values(identity).
		Returning("id", "identity_uid").
		LoadContext(ctx, &dbSbpAdmin)

	if err != nil {
		return app.SbpAdmin{}, err
	}

	return conversions.SbpAdminFromDB(dbSbpAdmin), tx.Commit()
}

func (s *Storage) GetOrCreateAdminIfNotExists(ctx context.Context, identity string) (app.SbpAdmin, error) {
	dbSbpAdmin, err := s.GetSbpAdmin(ctx, identity)

	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			return s.CreateSbpAdmin(ctx, identity)
		}

		return app.SbpAdmin{}, err
	}

	return dbSbpAdmin, err
}
