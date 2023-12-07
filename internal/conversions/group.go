package conversions

import (
	"fmt"
	"sbp/internal/entities"
	rabbitEntities "sbp/internal/entities/rabbit"

	uuid "github.com/satori/go.uuid"
)

func GroupFromRabbit(m rabbitEntities.ServerGroup) (entities.Group, error) {
	id, err := uuid.FromString(m.ID)
	if err != nil {
		return entities.Group{}, fmt.Errorf("unable to parse group ID: %w", entities.ErrBadRequest)
	}

	orgID, err := uuid.FromString(m.OrganizationID)
	if err != nil {
		return entities.Group{}, fmt.Errorf("unable to parse organization ID: %w", entities.ErrBadRequest)
	}

	return entities.Group{
		ID:             id,
		OrganizationID: orgID,
		Name:           m.Name,
		Description:    m.Description,
		IsDefault:      m.IsDefault,
		Deleted:        m.Deleted,
		Version:        m.Version,
	}, nil
}
