package payclientconverter

import (
	"fmt"
	logicEntities "sbp/internal/logic/entities"
	tinkoffEntities "sbp/internal/tinkoff/models"
)

// PaymentInitResponseToLogic ...
func PaymentInitResponseToLogic(response tinkoffEntities.ResponseInit) logicEntities.PaymentInit {
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

// GetQrResponseToLogic ...
func GetQrResponseToLogic(response tinkoffEntities.ResponseGetQr) logicEntities.PaymentGetQr {
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

// PaymentCancelResponseToLogic ...
func PaymentCancelResponseToLogic(response tinkoffEntities.ResponseCancel) logicEntities.PaymentCancel {
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
