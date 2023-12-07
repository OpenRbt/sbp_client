package conversions

import (
	"fmt"
	"sbp/internal/entities"
	rabbitEntities "sbp/internal/entities/rabbit"

	uuid "github.com/satori/go.uuid"
)

func UserFromRabbit(m rabbitEntities.AdminUser) (entities.User, error) {
	res := entities.User{
		ID:      m.ID,
		Email:   m.Email,
		Name:    m.Name,
		Role:    entities.Role(m.Role),
		Version: m.Version,
	}

	if m.OrganizationID != nil {
		orgID, err := uuid.FromString(*m.OrganizationID)
		if err != nil {
			return entities.User{}, fmt.Errorf("unable to parse organization ID: %w", entities.ErrBadRequest)
		}

		res.OrganizationID = &orgID
	}

	return res, nil
}
