package handlers

import (
	"fmt"
	conversions "sbp/internal/conversions"
	"sbp/internal/entities"
	"sbp/openapi/restapi/operations/washes"

	uuid "github.com/satori/go.uuid"
)

func (handler Handler) GetWashByID(params washes.GetWashByIDParams, auth *entities.Auth) washes.GetWashByIDResponder {
	op := "Get wash by ID:"
	resp := washes.NewGetWashByIDDefault(500)

	id, err := uuid.FromString(string(params.ID))
	if err != nil {
		err = fmt.Errorf("unable to parse wash id: %w", entities.ErrBadRequest)
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	res, err := handler.svc.GetWashByID(params.HTTPRequest.Context(), auth, id)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	return washes.NewGetWashByIDOK().WithPayload(conversions.WashToRest(res))
}

func (handler Handler) CreateWash(params washes.CreateWashParams, auth *entities.Auth) washes.CreateWashResponder {
	op := "Create wash:"
	resp := washes.NewCreateWashDefault(500)

	washCreationModel, err := conversions.WashCreationFromRest(*params.Body, auth.User.ID)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	newServer, err := handler.svc.CreateWash(params.HTTPRequest.Context(), auth, washCreationModel)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	return washes.NewCreateWashOK().WithPayload(conversions.WashToRest(newServer))
}

func (handler Handler) UpdateWash(params washes.UpdateWashParams, auth *entities.Auth) washes.UpdateWashResponder {
	op := "Update wash:"
	resp := washes.NewUpdateWashDefault(500)

	washID, err := uuid.FromString(string(params.ID))
	if err != nil {
		err = fmt.Errorf("unable to parse wash id: %w", entities.ErrBadRequest)
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	updateWashFromRest := conversions.WashUpdateFromRest(*params.Body)

	err = handler.svc.UpdateWash(params.HTTPRequest.Context(), auth, washID, updateWashFromRest)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	return washes.NewUpdateWashNoContent()
}

func (handler Handler) DeleteWash(params washes.DeleteWashParams, auth *entities.Auth) washes.DeleteWashResponder {
	op := "Delete wash:"
	resp := washes.NewDeleteWashDefault(500)

	washID, err := uuid.FromString(string(params.ID))
	if err != nil {
		err = fmt.Errorf("unable to parse wash id: %w", entities.ErrBadRequest)
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	err = handler.svc.DeleteWash(params.HTTPRequest.Context(), auth, washID)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	return washes.NewDeleteWashNoContent()
}

func (handler Handler) GetWashes(params washes.GetWashesParams, auth *entities.Auth) washes.GetWashesResponder {
	op := "Get washes:"
	resp := washes.NewGetWashesDefault(500)

	filter, err := conversions.WashFilterFromRest(params)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	res, err := handler.svc.GetWashes(params.HTTPRequest.Context(), auth, filter)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	payload := conversions.WashesToRest(res)
	return washes.NewGetWashesOK().WithPayload(payload)
}

func (handler Handler) AssignWashToGroup(params washes.AssignWashToGroupParams, auth *entities.Auth) washes.AssignWashToGroupResponder {
	op := "Assign wash to group:"
	resp := washes.NewAssignWashToGroupDefault(500)

	washID, err := uuid.FromString(string(params.WashID))
	if err != nil {
		err = fmt.Errorf("unable to parse wash id: %w", entities.ErrBadRequest)
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	groupID, err := uuid.FromString(string(params.GroupID))
	if err != nil {
		err = fmt.Errorf("unable to parse group id: %w", entities.ErrBadRequest)
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	err = handler.svc.AssignWashToGroup(params.HTTPRequest.Context(), auth, washID, groupID)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	return washes.NewAssignWashToGroupNoContent()
}
