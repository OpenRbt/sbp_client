package handlers

import (
	"log"
	restConverter "sbp/internal/api/rest/converter"
	logicEntities "sbp/internal/logic/entities"
	washServers "sbp/internal/openapi/restapi/operations/wash_servers"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// GetWashServer ...
func (handler Handler) GetWashServer(params washServers.GetWashServerParams, auth *logicEntities.Auth) washServers.GetWashServerResponder {
	id, err := uuid.FromString(params.ID)
	if err != nil {
		return washServers.NewGetWashServerBadRequest()
	}

	res, err := handler.logic.GetWashServer(params.HTTPRequest.Context(), id)
	switch {
	case err == nil:
		return washServers.NewGetWashServerOK().WithPayload(restConverter.СonvertWashServerToRest(res))
	case errors.Is(err, logicEntities.ErrNotFound):
		return washServers.NewGetWashServerNotFound()
	default:
		return washServers.NewGetWashServerInternalServerError()
	}
}

// CreateWashServer ...
func (handler Handler) CreateWashServer(params washServers.CreateParams, auth *logicEntities.Auth) washServers.CreateResponder {
	registerWashServerFromRest := restConverter.СonvertRegisterWashServerFromRest(*params.Body)

	admin, err := handler.logic.GetOrCreateAdminIfNotExists(params.HTTPRequest.Context(), auth)
	if err != nil {
		log.Println(err)
		if errors.Is(err, logicEntities.ErrNotFound) {
			return washServers.NewCreateBadRequest()
		} else {
			return washServers.NewCreateInternalServerError()
		}
	}

	newServer, err := handler.logic.CreateWashServer(params.HTTPRequest.Context(), *admin, registerWashServerFromRest)
	if err != nil {
		log.Println(err)
		if errors.Is(err, logicEntities.ErrNotFound) {
			return washServers.NewCreateBadRequest()
		} else {
			return washServers.NewCreateInternalServerError()
		}
	}

	return washServers.NewCreateOK().WithPayload(restConverter.СonvertWashServerToRest(newServer))
}

// UpdateWashServer ...
func (handler Handler) UpdateWashServer(params washServers.UpdateParams, auth *logicEntities.Auth) washServers.UpdateResponder {

	admin, err := handler.logic.GetOrCreateAdminIfNotExists(params.HTTPRequest.Context(), auth)
	if err != nil {
		log.Println(err)
		if errors.Is(err, logicEntities.ErrNotFound) {
			return washServers.NewUpdateBadRequest()
		} else {
			return washServers.NewUpdateInternalServerError()
		}
	}

	updateWashServerFromRest, err := restConverter.СonvertUpdateWashServerFromRest(*params.Body)
	if err != nil {
		return washServers.NewUpdateBadRequest()
	}

	err = handler.logic.UpdateWashServer(params.HTTPRequest.Context(), *admin, updateWashServerFromRest)
	switch {
	case err == nil:
		return washServers.NewUpdateNoContent()
	case errors.Is(err, logicEntities.ErrNotFound):
		return washServers.NewUpdateNotFound()
	default:
		return washServers.NewUpdateInternalServerError()
	}
}

// DeleteWashServer ...
func (handler Handler) DeleteWashServer(params washServers.DeleteParams, auth *logicEntities.Auth) washServers.DeleteResponder {

	admin, err := handler.logic.GetOrCreateAdminIfNotExists(params.HTTPRequest.Context(), auth)
	if err != nil {
		log.Println(err)
		if errors.Is(err, logicEntities.ErrNotFound) {
			return washServers.NewDeleteBadRequest()
		} else {
			return washServers.NewDeleteInternalServerError()
		}
	}

	ctx := params.HTTPRequest.Context()
	deleteWashServerFromRest, err := restConverter.СonvertDeleteWashServerFromRest(*params.Body)
	if err != nil {
		return washServers.NewDeleteBadRequest()
	}

	err = handler.logic.DeleteWashServer(ctx, *admin, deleteWashServerFromRest)
	switch {
	case err == nil:
		return washServers.NewDeleteNoContent()
	case errors.Is(err, logicEntities.ErrNotFound):
		return washServers.NewDeleteNotFound()
	default:
		return washServers.NewDeleteInternalServerError()
	}
}

// GetWashServerList ...
func (handler Handler) GetWashServerList(params washServers.ListParams, auth *logicEntities.Auth) washServers.ListResponder {
	ctx := params.HTTPRequest.Context()
	res, err := handler.logic.GetWashServerList(ctx, restConverter.СonvertPaginationFromRest(*params.Body))

	switch {
	case err == nil:
		payload := restConverter.СonvertWashServerListToRest(res)
		return washServers.NewListOK().WithPayload(payload)
	case len(res) == 0:
		return washServers.NewListNotFound()
	default:
		return washServers.NewListInternalServerError()
	}
}
