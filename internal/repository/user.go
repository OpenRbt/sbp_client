package repository

import (
	"context"
	"errors"
	"sbp/internal/entities"
	"sbp/internal/helpers"

	"github.com/gocraft/dbr/v2"
)

var UserColumns = []string{
	"id", "email", "name", "role", "organization_id", "version", "deleted",
}

func (s *Repository) GetUserByID(ctx context.Context, id string) (entities.User, error) {
	var user entities.User

	err := s.db.
		NewSession(nil).
		Select(UserColumns...).
		From("users").
		Where("id = ?", id).
		LoadOneContext(ctx, &user)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = entities.ErrNotFound
		}

		return entities.User{}, helpers.CustomError(layer, "GetUserByID", err)
	}

	return user, nil
}

func (s *Repository) InsertUser(ctx context.Context, user entities.User) error {
	userQuery, args, err := SQ.Insert("users").
		Columns(UserColumns...).
		Values(user.ID, user.Email, user.Name, user.Role, user.OrganizationID, user.Version, user.Deleted).
		Suffix(`
        	ON CONFLICT (id) DO UPDATE SET
            email = EXCLUDED.email,
            name = EXCLUDED.name, 
            role = EXCLUDED.role, 
            organization_id = EXCLUDED.organization_id, 
            version = EXCLUDED.version,
			deleted = EXCLUDED.deleted
		`).
		ToSql()

	if err != nil {
		return helpers.CustomError(layer, "InsertUser", err)
	}

	_, err = s.db.NewSession(nil).ExecContext(ctx, userQuery, args...)
	if err != nil {
		return helpers.CustomError(layer, "InsertUser", err)
	}

	return nil
}
