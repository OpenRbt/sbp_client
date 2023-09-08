package handlers

import (
	"log"
	restConverter "sbp/internal/api/rest/converter"
	logicEntities "sbp/internal/logic/entities"
	wash "sbp/openapi/restapi/operations/wash"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// GetWash ...
func (handler Handler) GetWash(params wash.GetWashParams, auth *logicEntities.AuthExtended) wash.GetWashResponder {
	id, err := uuid.FromString(params.ID)
	if err != nil {
		return wash.NewGetWashBadRequest()
	}

	res, err := handler.logic.GetWash(params.HTTPRequest.Context(), id)
	switch {
	case err == nil:
		return wash.NewGetWashOK().WithPayload(restConverter.СonvertWashToRest(res))
	case errors.Is(err, logicEntities.ErrNotFound):
		return wash.NewGetWashNotFound()
	default:
		return wash.NewGetWashInternalServerError()
	}
}

// CreateWash ...
func (handler Handler) CreateWash(params wash.CreateParams, auth *logicEntities.AuthExtended) wash.CreateResponder {
	registerWashFromRest := restConverter.СonvertRegisterWashFromRest(*params.Body)
	registerWashFromRest.OwnerID = auth.User.ID
	newServer, err := handler.logic.CreateWash(params.HTTPRequest.Context(), registerWashFromRest)
	if err != nil {
		log.Println(err)
		if errors.Is(err, logicEntities.ErrNotFound) {
			return wash.NewCreateBadRequest()
		} else {
			return wash.NewCreateInternalServerError()
		}
	}

	return wash.NewCreateOK().WithPayload(restConverter.СonvertWashToRest(newServer))
}

// UpdateWash ...
func (handler Handler) UpdateWash(params wash.UpdateParams, auth *logicEntities.AuthExtended) wash.UpdateResponder {
	updateWashFromRest, err := restConverter.СonvertUpdateWashFromRest(*params.Body)
	if err != nil {
		return wash.NewUpdateBadRequest()
	}
	updateWashFromRest.OwnerID = auth.User.ID
	err = handler.logic.UpdateWash(params.HTTPRequest.Context(), updateWashFromRest)
	switch {
	case err == nil:
		return wash.NewUpdateNoContent()
	case errors.Is(err, logicEntities.ErrNotFound):
		return wash.NewUpdateNotFound()
	default:
		return wash.NewUpdateInternalServerError()
	}
}

// DeleteWash ...
func (handler Handler) DeleteWash(params wash.DeleteParams, auth *logicEntities.AuthExtended) wash.DeleteResponder {
	ctx := params.HTTPRequest.Context()
	washId, err := restConverter.СonvertDeleteWashFromRest(*params.Body)
	if err != nil {
		return wash.NewDeleteBadRequest()
	}
	err = handler.logic.DeleteWash(ctx, auth.User.ID, washId)
	switch {
	case err == nil:
		return wash.NewDeleteNoContent()
	case errors.Is(err, logicEntities.ErrNotFound):
		return wash.NewDeleteNotFound()
	default:
		return wash.NewDeleteInternalServerError()
	}
}

// GetWashList ...
func (handler Handler) GetWashList(params wash.ListParams, auth *logicEntities.AuthExtended) wash.ListResponder {
	ctx := params.HTTPRequest.Context()
	res, err := handler.logic.GetWashList(ctx, restConverter.СonvertPaginationFromRest(*params.Body))

	switch {
	case err == nil:
		payload := restConverter.СonvertWashListToRest(res)
		return wash.NewListOK().WithPayload(payload)
	case len(res) == 0:
		return wash.NewListNotFound()
	default:
		return wash.NewListInternalServerError()
	}
}
