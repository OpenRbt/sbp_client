package handlers

import (
	conversions "sbp/internal/conversions"
	"sbp/internal/entities"
	openapiEntities "sbp/openapi/models"
	"sbp/openapi/restapi/operations/notifications"
	"sbp/openapi/restapi/operations/payments"
)

func (handler *Handler) CancelPayment(params payments.CancelPaymentParams, auth *entities.Auth) payments.CancelPaymentResponder {
	op := "Cancel payment:"
	resp := payments.NewCancelPaymentDefault(500)

	req := entities.PaymentСancellationRequest{
		WashID:  params.Body.WashID,
		PostID:  params.Body.PostID,
		OrderID: params.Body.OrderID,
	}

	_, err := handler.svc.CancelPayment(params.HTTPRequest.Context(), req)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	return payments.NewCancelPaymentOK()
}

func (handler *Handler) InitPayment(params payments.InitPaymentParams, auth *entities.Auth) payments.InitPaymentResponder {
	op := "Init payment:"
	resp := payments.NewInitPaymentDefault(500)

	req := entities.PaymentRequest{
		Amount:  params.Body.Amount,
		WashID:  params.Body.WashID,
		PostID:  params.Body.PostID,
		OrderID: params.Body.OrderID,
	}

	pay, err := handler.svc.InitPayment(params.HTTPRequest.Context(), req)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	payResp := openapiEntities.PaymentResponse{
		OrderID: pay.OrderID,
		URL:     pay.UrlPay,
	}
	return payments.NewInitPaymentOK().WithPayload(&payResp)
}

func (handler *Handler) ReceiveNotification(params notifications.ReceiveNotificationParams, auth *entities.Auth) notifications.ReceiveNotificationResponder {
	op := "Recieve notification:"
	resp := notifications.NewReceiveNotificationDefault(500)

	registerNotif := conversions.PaymentNotificationFromRest(*params.Body)
	err := handler.svc.ReceiveNotification(params.HTTPRequest.Context(), registerNotif)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	return notifications.NewReceiveNotificationOK().WithPayload("OK")
}
