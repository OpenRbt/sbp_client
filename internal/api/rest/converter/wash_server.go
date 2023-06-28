package restconverter

import (
	logicEntities "sbp/internal/logic/entities"
	openapiEntities "sbp/internal/openapi/models"

	uuid "github.com/satori/go.uuid"
)

// СonvertWashServerToRest ...
func СonvertWashServerToRest(washServer logicEntities.WashServer) *openapiEntities.WashServer {
	return &openapiEntities.WashServer{
		ID:               washServer.ID.String(),
		Name:             washServer.Title,
		Description:      washServer.Description,
		ServiceKey:       washServer.ServiceKey,
		TerminalPassword: washServer.TerminalPassword,
	}
}

// СonvertRegisterWashServerFromRest ...
func СonvertRegisterWashServerFromRest(rest openapiEntities.WashServerCreate) logicEntities.RegisterWashServer {
	return logicEntities.RegisterWashServer{
		Title:            *rest.Name,
		Description:      rest.Description,
		TerminalKey:      rest.TerminalKey,
		TerminalPassword: rest.TerminalPassword,
	}
}

// СonvertDeleteWashServerFromRest ...
func СonvertDeleteWashServerFromRest(model openapiEntities.WashServerDelete) (uuid.UUID, error) {
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

// СonvertWashServerListToRest ...
func СonvertWashServerListToRest(washServerEntity []logicEntities.WashServer) []*openapiEntities.WashServer {
	res := make([]*openapiEntities.WashServer, len(washServerEntity))

	for i, value := range washServerEntity {
		rest := СonvertWashServerToRest(value)
		res[i] = rest
	}

	return res
}

// СonvertUpdateWashServerFromRest ...
func СonvertUpdateWashServerFromRest(model openapiEntities.WashServerUpdate) (logicEntities.UpdateWashServer, error) {
	id, err := uuid.FromString(model.ID)
	if err != nil {
		return logicEntities.UpdateWashServer{}, err
	}

	return logicEntities.UpdateWashServer{
		ID:          id,
		Title:       &model.Name,
		Description: &model.Description,
	}, nil
}
