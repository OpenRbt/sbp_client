package repository

import (
	"context"
	"errors"
	"sbp/internal/entities"
	"sbp/internal/helpers"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

var WashColumns = []string{
	"id", "owner_id", "password", "title", "description", "terminal_key", "terminal_password", "created_at", "updated_at",
}

func (s *Repository) CreateWash(ctx context.Context, newWash entities.WashCreation) (entities.Wash, error) {
	var wash entities.Wash

	err := s.db.NewSession(nil).
		InsertInto("washes").
		Columns("owner_id", "password", "title", "description", "terminal_key", "terminal_password").
		Record(newWash).
		Returning("id", "owner_id", "password", "title", "description", "terminal_key", "terminal_password", "created_at", "updated_at").
		LoadContext(ctx, &wash)

	if err != nil {
		return entities.Wash{}, helpers.CustomError(layer, "CreateWash", err)
	}

	return wash, nil
}

func (s *Repository) GetWashByID(ctx context.Context, id uuid.UUID) (entities.Wash, error) {
	var wash entities.Wash

	err := s.db.NewSession(nil).
		Select("w.id, w.group_id, w.owner_id, w.password, w.title, w.description, w.terminal_key, w.terminal_password, w.created_at, w.updated_at, org.id organization_id").
		From(dbr.I("washes").As("w")).
		LeftJoin(dbr.I("wash_groups").As("gr"), "w.group_id = gr.id").
		LeftJoin(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("w.id = ? AND NOT w.deleted", id).
		LoadOneContext(ctx, &wash)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = entities.ErrNotFound
		}

		return entities.Wash{}, helpers.CustomError(layer, "GetWash", err)
	}

	return wash, nil
}

func (s *Repository) UpdateWash(ctx context.Context, id uuid.UUID, updateWash entities.WashUpdate) error {
	query := s.db.NewSession(nil).Update("washes").
		Where("id = ?", id)

	if updateWash.Title != nil {
		query = query.Set("title", updateWash.Title)
	}
	if updateWash.Description != nil {
		query = query.Set("description", updateWash.Description)
	}
	if updateWash.TerminalKey != nil {
		query = query.Set("terminal_key", updateWash.TerminalKey)
	}
	if updateWash.TerminalPassword != nil {
		query = query.Set("terminal_password", updateWash.TerminalPassword)
	}

	_, err := query.ExecContext(ctx)
	if err != nil {
		return helpers.CustomError(layer, "UpdateWash", err)
	}

	return nil
}

func (s *Repository) DeleteWash(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewSession(nil).
		Update("washes").
		Where("id = ? AND NOT DELETED", id).
		Set("deleted", true).
		ExecContext(ctx)
	if err != nil {
		return helpers.CustomError(layer, "DeleteWash", err)
	}

	return nil
}

func (s *Repository) GetWashes(ctx context.Context, filter entities.WashFilter) ([]entities.Wash, error) {
	var washes []entities.Wash

	query := s.db.NewSession(nil).
		Select("w.id, w.group_id, w.owner_id, w.password, w.title, w.description, w.terminal_key, w.terminal_password, w.created_at, w.updated_at, org.id organization_id").
		From(dbr.I("washes").As("w")).
		LeftJoin(dbr.I("wash_groups").As("gr"), "w.group_id = gr.id").
		LeftJoin(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("NOT w.deleted")

	if filter.GroupID != nil {
		query = query.Where("w.group_id = ?", filter.GroupID)
	}

	if filter.OrganizationID != nil {
		query = query.Where("org.id = ?", filter.OrganizationID)
	}

	count, err := query.Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		LoadContext(ctx, &washes)

	if err != nil {
		return nil, helpers.CustomError(layer, "GetWashes", err)
	}

	if count == 0 {
		return nil, nil
	}

	return washes, nil
}

func (s *Repository) AssignWashToGroup(ctx context.Context, washID, groupID uuid.UUID) error {
	_, err := s.db.NewSession(nil).
		Update("washes").
		Where("id = ? AND NOT deleted", washID).
		Set("group_id", groupID).
		ExecContext(ctx)
	if err != nil {
		return helpers.CustomError(layer, "AssignWashToGroup", err)
	}

	return nil
}
