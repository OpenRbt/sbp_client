package repository

import (
	"context"
	"errors"
	"sbp/internal/entities"
	"sbp/internal/helpers"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

var GroupColumns = []string{
	"id", "name", "description", "organization_id", "is_default", "deleted", "version",
}

func (r *Repository) GetGroupByID(ctx context.Context, id uuid.UUID) (entities.Group, error) {
	op := "GetGroupByID"

	var group entities.Group
	err := r.db.NewSession(nil).
		Select(GroupColumns...).
		From("wash_groups").
		Where("id = ? AND NOT deleted", id).
		LoadOneContext(ctx, &group)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = entities.ErrNotFound
		}
		return entities.Group{}, helpers.CustomError(layer, op, err)
	}

	return group, nil
}

func (r *Repository) InsertGroup(ctx context.Context, group entities.Group) error {
	op := "InsertGroup"

	groupQuery, args, err := SQ.Insert("wash_groups").
		Columns(GroupColumns...).
		Values(group.ID, group.Name, group.Description, group.OrganizationID, group.IsDefault, group.Deleted, group.Version).
		Suffix(`
			ON CONFLICT (id) DO UPDATE SET
				name = EXCLUDED.name,
				description = EXCLUDED.description,
				organization_id = EXCLUDED.organization_id,
				is_default = EXCLUDED.is_default,
				deleted = EXCLUDED.deleted,
				version = EXCLUDED.version
		`).
		ToSql()
	if err != nil {
		return helpers.CustomError(layer, op, err)
	}

	_, err = r.db.NewSession(nil).ExecContext(ctx, groupQuery, args...)
	if err != nil {
		return helpers.CustomError(layer, op, err)
	}

	return nil
}
