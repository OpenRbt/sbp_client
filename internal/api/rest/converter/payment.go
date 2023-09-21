package restconverter

import (
	logicEntities "sbp/internal/logic/entities"
	openapiEntities "sbp/openapi/models"
)

// СonvertRegisterNotificationFromRest ...
func СonvertRegisterNotificationFromRest(req openapiEntities.Notification) logicEntities.PaymentNotification {
	return logicEntities.PaymentNotification{
		Success:     req.Success,
		Amount:      req.Amount,
		ErrorCode:   req.ErrorCode,
		OrderID:     req.OrderID,
		Pan:         req.Pan,
		PaymentID:   req.PaymentID,
		Status:      req.Status,
		TerminalKey: req.TerminalKey,
		Token:       req.Token,
	}
}
