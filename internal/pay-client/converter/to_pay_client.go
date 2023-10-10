package payclientconverter

import (
	logicEntities "sbp/internal/logic/entities"
	tinkoffEntities "sbp/internal/tinkoff/models"
)

// PaymentNotificationToPayClient ...
func PaymentNotificationToPayClient(req logicEntities.PaymentNotification) tinkoffEntities.Notification {
	return tinkoffEntities.Notification{
		Amount:      req.Amount,
		ErrorCode:   req.ErrorCode,
		OrderID:     req.OrderID,
		Pan:         req.Pan,
		PaymentID:   req.PaymentID,
		Status:      req.Status,
		Success:     req.Success,
		TerminalKey: req.TerminalKey,
		Token:       req.Token,
	}
}
