package repository

import (
	"context"
	"errors"
	"sbp/internal/entities"
	"sbp/internal/helpers"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

var OrganizationColumns = []string{
	"id", "name", "display_name", "description", "is_default", "deleted", "version",
}

func (r *Repository) GetOrganizationByID(ctx context.Context, id uuid.UUID) (entities.Organization, error) {
	op := "GetOrganizationByID"

	var org entities.Organization
	err := r.db.NewSession(nil).
		Select(OrganizationColumns...).
		From("organizations").
		Where("id = ? AND NOT deleted", id).
		LoadOneContext(ctx, &org)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = entities.ErrNotFound
		}
		return entities.Organization{}, helpers.CustomError(layer, op, err)
	}

	return org, nil
}

func (r *Repository) InsertOrganization(ctx context.Context, org entities.Organization) error {
	op := "InsertOrganization"

	organizationQuery, args, err := SQ.Insert("organizations").
		Columns("id", "name", "display_name", "description", "is_default", "deleted", "version").
		Values(org.ID, org.Name, org.DisplayName, org.Description, org.IsDefault, org.Deleted, org.Version).
		Suffix(`
			ON CONFLICT (id) DO UPDATE SET
				name = EXCLUDED.name,
				display_name = EXCLUDED.display_name,
				description = EXCLUDED.description,
				is_default = EXCLUDED.is_default,
				deleted = EXCLUDED.deleted,
				version = EXCLUDED.version
		`).
		ToSql()
	if err != nil {
		return helpers.CustomError(layer, op, err)
	}

	_, err = r.db.NewSession(nil).ExecContext(ctx, organizationQuery, args...)
	if err != nil {
		return helpers.CustomError(layer, op, err)
	}

	return nil
}
