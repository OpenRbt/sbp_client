package conversions

import (
	"fmt"
	"sbp/internal/entities"
	rabbitEntities "sbp/internal/entities/rabbit"

	uuid "github.com/satori/go.uuid"
)

func OrganizationFromRabbit(m rabbitEntities.Organization) (entities.Organization, error) {
	id, err := uuid.FromString(m.ID)
	if err != nil {
		return entities.Organization{}, fmt.Errorf("unable to parse organization ID: %w", entities.ErrBadRequest)
	}

	return entities.Organization{
		ID:          id,
		Name:        m.Name,
		DisplayName: m.DisplayName,
		Description: m.Description,
		IsDefault:   m.IsDefault,
		Deleted:     m.Deleted,
		Version:     m.Version,
	}, nil
}
