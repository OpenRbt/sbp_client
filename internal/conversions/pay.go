package conversions

import (
	"fmt"
	"sbp/internal/app"
	"sbp/internal/transport/pay/models"
)

func InitFromResponse(response models.ResponseInit) app.Init {
	return app.Init{
		Success:   response.Success,
		OrderID:   response.OrderID,
		PaymentID: response.PaymentID,
		Status:    response.Status,
		Url:       response.PaymentURL,
	}
}

func GetQrFromResponse(response models.ResponseGetQr) app.GetQr {

	return app.GetQr{
		Success:   response.Success,
		ErrorCode: response.ErrorCode,
		Message:   response.Message,
		OrderID:   response.OrderID,
		PaymentID: fmt.Sprint(response.PaymentID),
		UrlPay:    response.Data,
	}
}

func CancelFromResponse(response models.ResponseCancel) app.Cancel {
	return app.Cancel{
		Success:   response.Success,
		Status:    response.Status,
		ErrorCode: response.ErrorCode,
		OrderID:   response.OrderID,
		PaymentID: fmt.Sprint(response.PaymentID),
	}
}
