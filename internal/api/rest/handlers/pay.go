package handlers

import (
	restConverter "sbp/internal/api/rest/converter"
	logicEntities "sbp/internal/logic/entities"

	openapiEntities "sbp/openapi/models"
	washes "sbp/openapi/restapi/operations/wash"
)

// Cancel ...
func (handler *Handler) Cancel(params washes.CancelParams, auth *logicEntities.AuthExtended) washes.CancelResponder {
	// auth
	if !auth.IsAdmin() {
		return washes.NewCancelForbidden().WithPayload(&openapiEntities.Error{
			Code:    &ErrAccessDeniedCode,
			Message: &ErrAccessDenied,
		})
	}
	//
	req := logicEntities.PayСancellationRequest{
		WashID:  params.Body.WashID,
		PostID:  params.Body.PostID,
		OrderID: params.Body.OrderID,
	}
	_, err := handler.logic.Cancel(params.HTTPRequest.Context(), req)
	switch {
	case err == nil:
		return washes.NewCancelOK()
	default:
		handler.logger.Error(err)
		return washes.NewCancelBadRequest()
	}
}

// Pay ...
func (handler *Handler) Pay(params washes.PayParams, auth *logicEntities.AuthExtended) washes.PayResponder {
	// auth
	if !auth.IsAdmin() {
		return washes.NewPayForbidden().WithPayload(&openapiEntities.Error{
			Code:    &ErrAccessDeniedCode,
			Message: &ErrAccessDenied,
		})
	}
	//
	//
	req := logicEntities.PayRequest{
		Amount:  params.Body.Amount,
		WashID:  params.Body.WashID,
		PostID:  params.Body.PostID,
		OrderID: params.Body.OrderID,
	}
	resp, err := handler.logic.Pay(params.HTTPRequest.Context(), req)
	switch {
	case err == nil:
		payResp := openapiEntities.PayResponse{
			OrderID: resp.OrderID,
			URL:     resp.UrlPay,
		}
		return washes.NewPayOK().WithPayload(&payResp)
	default:
		errCode := int32(0)
		message := err.Error()
		handler.logger.Error(err)
		return washes.NewPayBadRequest().WithPayload(&openapiEntities.Error{
			Code:    &errCode,
			Message: &message,
		})
	}
}

// Notif ...
func (handler *Handler) Notif(params washes.NotificationParams, auth *logicEntities.AuthExtended) washes.NotificationResponder {
	registerNotif := restConverter.СonvertRegisterNotificationFromRest(*params.Body)
	err := handler.logic.Notification(params.HTTPRequest.Context(), registerNotif)
	if err != nil {
		handler.logger.Error(err)
		return washes.NewNotificationInternalServerError()
	}

	return washes.NewNotificationOK().WithPayload("OK")
}
