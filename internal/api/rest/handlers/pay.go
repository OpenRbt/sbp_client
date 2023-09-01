package handlers

import (
	restConverter "sbp/internal/api/rest/converter"
	logicEntities "sbp/internal/logic/entities"
	openapiEntities "sbp/openapi/models"
	washServers "sbp/openapi/restapi/operations/wash_servers"
)

// Cancel ...
func (handler *Handler) Cancel(params washServers.CancelParams, auth *logicEntities.Auth) washServers.CancelResponder {
	req := logicEntities.PayСancellationRequest{
		OrderID: params.Body.OrderID,
	}
	_, err := handler.logic.Cancel(params.HTTPRequest.Context(), req)
	switch {
	case err == nil:
		return washServers.NewCancelOK()
	default:
		handler.logger.Error(err)
		return washServers.NewCancelBadRequest()
	}
}

// Pay ...
func (handler *Handler) Pay(params washServers.PayParams, auth *logicEntities.Auth) washServers.PayResponder {
	req := logicEntities.PayRequest{
		Amount:   params.Body.Amount,
		ServerID: params.Body.ServerID,
		PostID:   params.Body.PostID,
	}
	resp, err := handler.logic.Pay(params.HTTPRequest.Context(), req)
	switch {
	case err == nil:
		payResp := openapiEntities.PayResponse{
			URL:     resp.UrlPay,
			OrderID: resp.OrderID,
		}
		return washServers.NewPayOK().WithPayload(&payResp)
	default:
		errCode := int32(0)
		message := err.Error()
		handler.logger.Error(err)
		return washServers.NewPayBadRequest().WithPayload(&openapiEntities.Error{
			Code:    &errCode,
			Message: &message,
		})
	}
}

// Notif ...
func (handler *Handler) Notif(params washServers.NotificationParams, auth *logicEntities.Auth) washServers.NotificationResponder {
	registerNotif := restConverter.СonvertRegisterNotificationFromRest(*params.Body)
	err := handler.logic.Notification(params.HTTPRequest.Context(), registerNotif)
	if err != nil {
		handler.logger.Error(err)
		return washServers.NewNotificationInternalServerError()
	}

	return washServers.NewNotificationOK().WithPayload("OK")
}
