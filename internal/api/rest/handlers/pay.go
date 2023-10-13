package handlers

import (
	restConverter "sbp/internal/api/rest/converter"
	logicEntities "sbp/internal/logic/entities"
	"strings"

	openapiEntities "sbp/openapi/models"
	washes "sbp/openapi/restapi/operations/wash"
)

// Cancel ...
func (handler *Handler) Cancel(params washes.CancelParams, auth *logicEntities.AuthExtended) washes.CancelResponder {
	// auth
	if !auth.IsAdmin() {
		handler.logger.Errorf("auth failed: user '%s' is not admin", auth.User.ID.String())
		return washes.NewCancelForbidden().WithPayload(&openapiEntities.Error{
			Code:    &ErrAccessDeniedCode,
			Message: &ErrAccessDenied,
		})
	}
	//
	req := logicEntities.PaymentСancellationRequest{
		WashID:  params.Body.WashID,
		PostID:  params.Body.PostID,
		OrderID: params.Body.OrderID,
	}
	_, err := handler.logic.Cancel(params.HTTPRequest.Context(), req)
	switch {
	case err == nil:
		return washes.NewCancelOK()
	default:
		handler.logger.Errorf("cancel payment failed: %w", err)
		return washes.NewCancelBadRequest()
	}
}

// Pay ...
func (handler *Handler) Pay(params washes.PayParams, auth *logicEntities.AuthExtended) washes.PayResponder {
	// auth
	if !auth.IsAdmin() {
		handler.logger.Errorf("auth failed: user '%s' is not admin", auth.User.ID.String())
		return washes.NewPayForbidden().WithPayload(&openapiEntities.Error{
			Code:    &ErrAccessDeniedCode,
			Message: &ErrAccessDenied,
		})
	}
	//
	//
	req := logicEntities.PaymentRequest{
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
		handler.logger.Errorf("payment request failed: %w", err)
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
		handler.logger.Errorf("notify request failed: %w", err)
		if strings.HasSuffix(registerNotif.TerminalKey, "DEMO") {
			return washes.NewNotificationOK().WithPayload("OK")
		}
		return washes.NewNotificationBadRequest().WithPayload("bad request")
	}

	return washes.NewNotificationOK().WithPayload("OK")
}
