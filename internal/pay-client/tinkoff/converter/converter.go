package tinkoffconverter

import (
	"fmt"
	logicEntities "sbp/internal/logic/entities"
	tinkoffEntities "sbp/internal/pay-client/tinkoff/models"
)

// ConvertInitFromResponse ...
func ConvertInitFromResponse(response tinkoffEntities.ResponseInit) logicEntities.PaymentInit {
	return logicEntities.PaymentInit{
		PaymentInfo: logicEntities.PaymentInfo{
			Success:   response.Success,
			OrderID:   response.OrderID,
			PaymentID: response.PaymentID,
		},
		Status: response.Status,
		Url:    response.PaymentURL,
	}
}

// ConvertGetQrFromResponse ...
func ConvertGetQrFromResponse(response tinkoffEntities.ResponseGetQr) logicEntities.PaymentGetQr {
	return logicEntities.PaymentGetQr{
		PaymentInfo: logicEntities.PaymentInfo{
			Success:   response.Success,
			OrderID:   response.OrderID,
			PaymentID: fmt.Sprint(response.PaymentID),
		},
		ErrorCode: response.ErrorCode,
		Message:   response.Message,
		UrlPay:    response.Data,
	}
}

// ConvertCancelFromResponse ...
func ConvertCancelFromResponse(response tinkoffEntities.ResponseCancel) logicEntities.PaymentCancel {
	return logicEntities.PaymentCancel{
		PaymentInfo: logicEntities.PaymentInfo{
			Success:   response.Success,
			OrderID:   response.OrderID,
			PaymentID: fmt.Sprint(response.PaymentID),
		},
		Status:    response.Status,
		ErrorCode: response.ErrorCode,
	}
}

// ConvertNotificationFromRequest ...
func ConvertNotificationFromRequest(req logicEntities.PaymentRegisterNotification) tinkoffEntities.Notification {
	return tinkoffEntities.Notification{
		Success:     req.Success,
		Amount:      int64(req.Amount),
		CardID:      int64(req.CardId),
		PaymentID:   req.PaymentID,
		OrderID:     req.OrderID,
		TerminalKey: req.TerminalKey,
		Status:      req.Status,
		ErrorCode:   req.ErrorCode,
		ExpDate:     req.ExpDate,
		Pan:         req.Pan,
	}
}
