package restconverter

import (
	logicEntities "sbp/internal/logic/entities"
	openapiEntities "sbp/openapi/models"
)

// СonvertRegisterNotificationFromRest ...
func СonvertRegisterNotificationFromRest(req openapiEntities.Notification) logicEntities.PaymentNotification {
	return logicEntities.PaymentNotification{
		AccountToken:     req.AccountToken,
		BankMemberID:     req.BankMemberID,
		BankMemberName:   req.BankMemberName,
		Details:          req.Details,
		ErrorCode:        req.ErrorCode,
		ExpDate:          req.ExpDate,
		Message:          req.Message,
		NotificationType: req.NotificationType,
		OrderID:          req.OrderID,
		Pan:              req.Pan,
		PaymentID:        req.PaymentID,
		PaymentURL:       req.PaymentURL,
		RebillID:         req.RebillID,
		RequestKey:       req.RequestKey,
		Status:           req.Status,
		TerminalKey:      req.TerminalKey,
		Token:            req.Token,
		Amount:           req.Amount,
		CardID:           req.CardID,
		Success:          req.Success,
	}
}
