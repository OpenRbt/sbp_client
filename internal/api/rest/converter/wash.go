package restconverter

import (
	logicEntities "sbp/internal/logic/entities"
	openapiEntities "sbp/openapi/models"

	uuid "github.com/satori/go.uuid"
)

// СonvertWashToRest ...
func СonvertWashToRest(washServer logicEntities.Wash) *openapiEntities.Wash {
	return &openapiEntities.Wash{
		Description:      washServer.Description,
		ID:               washServer.ID.String(),
		Name:             washServer.Title,
		Password:         washServer.Password,
		TerminalKey:      washServer.TerminalKey,
		TerminalPassword: washServer.TerminalPassword,
	}
}

// СonvertRegisterWashFromRest ...
func СonvertRegisterWashFromRest(rest openapiEntities.WashCreate) logicEntities.RegisterWash {
	return logicEntities.RegisterWash{
		Title:            *rest.Name,
		Description:      rest.Description,
		TerminalKey:      rest.TerminalKey,
		TerminalPassword: rest.TerminalPassword,
	}
}

// СonvertDeleteWashFromRest ...
func СonvertDeleteWashFromRest(model openapiEntities.WashDelete) (uuid.UUID, error) {
	id, err := uuid.FromString(*model.ID)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

// СonvertPaginationFromRest ...
func СonvertPaginationFromRest(model openapiEntities.Pagination) logicEntities.Pagination {
	return logicEntities.Pagination{
		Limit:  model.Limit,
		Offset: model.Offset,
	}
}

// СonvertWashListToRest ...
func СonvertWashListToRest(washServerEntity []logicEntities.Wash) []*openapiEntities.Wash {
	res := make([]*openapiEntities.Wash, len(washServerEntity))

	for i, value := range washServerEntity {
		rest := СonvertWashToRest(value)
		res[i] = rest
	}

	return res
}

// СonvertUpdateWashFromRest ...
func СonvertUpdateWashFromRest(model openapiEntities.WashUpdate) (logicEntities.UpdateWash, error) {
	id, err := uuid.FromString(model.ID)
	if err != nil {
		return logicEntities.UpdateWash{}, err
	}

	return logicEntities.UpdateWash{
		ID:          id,
		Title:       &model.Name,
		Description: &model.Description,
	}, nil
}
