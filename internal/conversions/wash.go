package conversions

import (
	"fmt"
	"sbp/internal/entities"
	"sbp/openapi/models"
	"sbp/openapi/restapi/operations/washes"

	"github.com/OpenRbt/share_business/wash_rabbit/entity/admin"
	"github.com/go-openapi/strfmt"

	uuid "github.com/satori/go.uuid"
)

func WashServerToRabbit(e entities.Wash) admin.ServerRegistered {
	return admin.ServerRegistered{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
	}
}

func WashServerUpdateToRabbit(id uuid.UUID, e entities.WashUpdate, deleted bool) admin.ServerUpdate {
	var del *bool
	if deleted {
		t := true
		del = &t
	}

	return admin.ServerUpdate{
		ID:          id.String(),
		Title:       e.Title,
		Description: e.Description,
		Deleted:     del,
	}
}

func WashToRest(washServer entities.Wash) *models.Wash {
	return &models.Wash{
		Description:      washServer.Description,
		ID:               washServer.ID.String(),
		Name:             washServer.Title,
		Password:         washServer.Password,
		TerminalKey:      washServer.TerminalKey,
		TerminalPassword: washServer.TerminalPassword,
		OrganizationID:   strfmt.UUID(washServer.OrganizationID.UUID.String()),
		GroupID:          strfmt.UUID(washServer.GroupID.UUID.String()),
	}
}

func WashCreationFromRest(rest models.WashCreation, ownerID string) (entities.WashCreation, error) {
	groupID, err := uuid.FromString(rest.GroupID.String())
	if err != nil {
		return entities.WashCreation{}, fmt.Errorf("wrong group ID: %w", entities.ErrBadRequest)
	}

	return entities.WashCreation{
		Title:            *rest.Name,
		Description:      *rest.Description,
		TerminalKey:      *rest.TerminalKey,
		TerminalPassword: *rest.TerminalPassword,
		GroupID:          groupID,
		OwnerID:          ownerID,
	}, nil
}

func PaginationFromRest(limit, offset int64) entities.Pagination {
	return entities.Pagination{
		Limit:  limit,
		Offset: offset,
	}
}

func WashFilterFromRest(params washes.GetWashesParams) (entities.WashFilter, error) {
	res := entities.WashFilter{
		Pagination: entities.Pagination{
			Limit:  *params.Limit,
			Offset: *params.Offset,
		},
	}

	if params.GroupID != nil {
		groupID, err := uuid.FromString(params.GroupID.String())
		if err != nil {
			return entities.WashFilter{}, fmt.Errorf("unable to parse groupID: %w", entities.ErrBadRequest)
		}

		res.GroupID = &groupID
	}

	return res, nil
}

func WashesToRest(washServerEntity []entities.Wash) []*models.Wash {
	res := make([]*models.Wash, len(washServerEntity))

	for i, value := range washServerEntity {
		rest := WashToRest(value)
		res[i] = rest
	}

	return res
}

func WashUpdateFromRest(model models.WashUpdate) entities.WashUpdate {
	return entities.WashUpdate{
		Title:            &model.Name,
		Description:      &model.Description,
		TerminalKey:      &model.TerminalKey,
		TerminalPassword: &model.TerminalPassword,
	}
}
