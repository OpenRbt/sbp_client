package repository

import (
	"context"
	"errors"
	logicEntities "sbp/internal/logic/entities"
	repConverter "sbp/internal/repository/converter"
	repEntities "sbp/internal/repository/entities"
	"sbp/pkg/bootstrap"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

// CreateWash ...
func (s *Repository) CreateWash(ctx context.Context, newWash logicEntities.RegisterWash) (logicEntities.Wash, error) {
	var registredServer repEntities.Wash

	err := s.db.NewSession(nil).
		InsertInto("washes").
		Columns("owner_id", "password", "title", "description", "terminal_key", "terminal_password").
		Record(repEntities.RegisterWash{
			OwnerID:          newWash.OwnerID,
			Password:         newWash.Password,
			Title:            newWash.Title,
			Description:      newWash.Description,
			TerminalKey:      newWash.TerminalKey,
			TerminalPassword: newWash.TerminalPassword,
		}).Returning("id", "owner_id", "password", "title", "description", "terminal_key", "terminal_password", "created_at", "updated_at").
		LoadContext(ctx, &registredServer)

	if err != nil {
		return logicEntities.Wash{}, bootstrap.CustomError(layer, "CreateWash", err)
	}

	return repConverter.ConvertWashFromDB(registredServer), bootstrap.CustomError(layer, "CreateWash", err)
}

// GetWash ...
func (s *Repository) GetWash(ctx context.Context, id uuid.UUID) (logicEntities.Wash, error) {
	var dbWash repEntities.Wash

	err := s.db.NewSession(nil).
		Select("*").
		From("washes").
		Where("id = ? AND NOT deleted", uuid.NullUUID{UUID: id, Valid: true}).
		LoadOneContext(ctx, &dbWash)

	switch {
	case err == nil:
		return repConverter.ConvertWashFromDB(dbWash), nil
	case errors.Is(err, dbr.ErrNotFound):
		return logicEntities.Wash{}, bootstrap.CustomError(layer, "GetWash", logicEntities.ErrNotFound)
	default:
		return logicEntities.Wash{}, bootstrap.CustomError(layer, "GetWash", err)
	}
}

// UpdateWash ...
func (s *Repository) UpdateWash(ctx context.Context, updateWash logicEntities.UpdateWash) error {
	dbUpdateWash := repConverter.ConvertUpdateWashToDb(updateWash)

	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return bootstrap.CustomError(layer, "UpdateWash", err)
	}

	updateStatement := tx.
		Update("washes").
		Where("id = ?", dbUpdateWash.ID)

	if dbUpdateWash.Name != nil {
		updateStatement = updateStatement.Set("title", dbUpdateWash.Name)
	}
	if dbUpdateWash.Description != nil {
		updateStatement = updateStatement.Set("description", dbUpdateWash.Description)
	}
	if dbUpdateWash.TerminalKey != nil {
		updateStatement = updateStatement.Set("terminal_key", dbUpdateWash.TerminalKey)
	}
	if dbUpdateWash.TerminalPassword != nil {
		updateStatement = updateStatement.Set("terminal_password", dbUpdateWash.TerminalPassword)
	}

	_, err = updateStatement.ExecContext(ctx)

	if err != nil {
		return bootstrap.CustomError(layer, "UpdateWash", err)
	}

	return tx.Commit()
}

// DeleteWash ...
func (s *Repository) DeleteWash(ctx context.Context, id uuid.UUID) error {
	dbDeleteWash := repConverter.ConvertDeleteWashToDB(id)

	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return bootstrap.CustomError(layer, "DeleteWash", err)
	}

	deleteStatement := tx.
		Update("washes").
		Where("id = ? AND NOT DELETED", dbDeleteWash.ID).
		Set("deleted", true)

	_, err = deleteStatement.ExecContext(ctx)

	if err != nil {
		return bootstrap.CustomError(layer, "DeleteWash", err)
	}

	return tx.Commit()
}

// GetWashList ...
func (s *Repository) GetWashList(ctx context.Context, pagination logicEntities.Pagination) ([]logicEntities.Wash, error) {
	var dbWashList []repEntities.Wash

	count, err := s.db.NewSession(nil).
		Select("*").
		From("washes").
		Where("NOT DELETED").
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &dbWashList)

	if err != nil {
		return []logicEntities.Wash{}, bootstrap.CustomError(layer, "GetWashList", err)
	}

	if count == 0 {
		return []logicEntities.Wash{}, nil
	}

	washServerListFromDB := repConverter.ConvertWashListFromDB(dbWashList)

	return washServerListFromDB, nil
}
