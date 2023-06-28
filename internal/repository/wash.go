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

// CreateWashServer ...
func (s *Repository) CreateWashServer(ctx context.Context, admin uuid.UUID, newWashServer logicEntities.RegisterWashServer) (logicEntities.WashServer, error) {
	var registredServer repEntities.WashServer

	err := s.db.NewSession(nil).
		InsertInto("wash_servers").
		Columns("title", "description", "owner", "service_key", "terminal_key", "terminal_password").
		Record(repEntities.RegisterWashServer{
			Title:       newWashServer.Title,
			Description: newWashServer.Description,
			Owner: uuid.NullUUID{
				UUID:  admin,
				Valid: true,
			},
			ServiceKey:       s.generateNewServiceKey(),
			TerminalKey:      newWashServer.TerminalKey,
			TerminalPassword: newWashServer.TerminalPassword,
		}).Returning("id", "title", "description", "owner", "service_key", "terminal_key", "terminal_password").
		LoadContext(ctx, &registredServer)

	if err != nil {
		return logicEntities.WashServer{}, bootstrap.CustomError(layer, "CreateWashServer", err)
	}

	return repConverter.ConvertWashServerFromDB(registredServer), bootstrap.CustomError(layer, "CreateWashServer", err)
}

// GetWashServer ...
func (s *Repository) GetWashServer(ctx context.Context, id uuid.UUID) (logicEntities.WashServer, error) {
	var dbWashServer repEntities.WashServer

	err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("id = ? AND NOT deleted", uuid.NullUUID{UUID: id, Valid: true}).
		LoadOneContext(ctx, &dbWashServer)

	switch {
	case err == nil:
		return repConverter.ConvertWashServerFromDB(dbWashServer), nil
	case errors.Is(err, dbr.ErrNotFound):
		return logicEntities.WashServer{}, bootstrap.CustomError(layer, "GetWashServer", logicEntities.ErrNotFound)
	default:
		return logicEntities.WashServer{}, bootstrap.CustomError(layer, "GetWashServer", err)
	}
}

// UpdateWashServer ...
func (s *Repository) UpdateWashServer(ctx context.Context, updateWashServer logicEntities.UpdateWashServer) error {
	dbUpdateWashServer := repConverter.ConvertUpdateWashServerToDb(updateWashServer)

	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return bootstrap.CustomError(layer, "UpdateWashServer", err)
	}

	updateStatement := tx.
		Update("wash_servers").
		Where("id = ?", dbUpdateWashServer.ID)

	if dbUpdateWashServer.Name != nil {
		updateStatement = updateStatement.Set("title", dbUpdateWashServer.Name)
	}
	if dbUpdateWashServer.Description != nil {
		updateStatement = updateStatement.Set("description", dbUpdateWashServer.Description)
	}
	if dbUpdateWashServer.TerminalKey != nil {
		updateStatement = updateStatement.Set("terminal_key", dbUpdateWashServer.TerminalKey)
	}
	if dbUpdateWashServer.TerminalPassword != nil {
		updateStatement = updateStatement.Set("terminal_password", dbUpdateWashServer.TerminalPassword)
	}

	_, err = updateStatement.ExecContext(ctx)

	if err != nil {
		return bootstrap.CustomError(layer, "UpdateWashServer", err)
	}

	return tx.Commit()
}

// DeleteWashServer ...
func (s *Repository) DeleteWashServer(ctx context.Context, id uuid.UUID) error {
	dbDeleteWashServer := repConverter.ConvertDeleteWashServerToDB(id)

	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return bootstrap.CustomError(layer, "DeleteWashServer", err)
	}

	deleteStatement := tx.
		Update("wash_servers").
		Where("id = ? AND NOT DELETED", dbDeleteWashServer.ID).
		Set("deleted", true)

	_, err = deleteStatement.ExecContext(ctx)

	if err != nil {
		return bootstrap.CustomError(layer, "DeleteWashServer", err)
	}

	return tx.Commit()
}

// GetWashServerList ...
func (s *Repository) GetWashServerList(ctx context.Context, pagination logicEntities.Pagination) ([]logicEntities.WashServer, error) {
	var dbWashServerList []repEntities.WashServer

	count, err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("NOT DELETED").
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &dbWashServerList)

	if err != nil {
		return []logicEntities.WashServer{}, bootstrap.CustomError(layer, "GetWashServerList", err)
	}

	if count == 0 {
		return []logicEntities.WashServer{}, nil
	}

	washServerListFromDB := repConverter.ConvertWashServerListFromDB(dbWashServerList)

	return washServerListFromDB, nil
}
