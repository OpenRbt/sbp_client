package conversions

import (
	"fmt"
	"sbp/internal/entities"
	rabbitEntities "sbp/internal/entities/rabbit"
	apiModels "sbp/openapi/models"
	payModels "sbp/tinkoffapi/models"
)

func InitPaymentFromRest(response payModels.ResponseInit) entities.PaymentInit {
	return entities.PaymentInit{
		PaymentInfo: entities.PaymentInfo{
			Success:   response.Success,
			OrderID:   response.OrderID,
			PaymentID: response.PaymentID,
		},
		Status: response.Status,
		Url:    response.PaymentURL,
	}
}

func GetQRFromRest(resp payModels.ResponseGetQr) entities.PaymentGetQr {

	return entities.PaymentGetQr{
		PaymentInfo: entities.PaymentInfo{
			Success:   resp.Success,
			OrderID:   resp.OrderID,
			PaymentID: fmt.Sprint(resp.PaymentID),
		},
		ErrorCode: resp.ErrorCode,
		Message:   resp.Message,
		UrlPay:    resp.Data,
	}
}

func CancelPaymentFromRest(resp payModels.ResponseCancel) entities.PaymentCancel {
	return entities.PaymentCancel{
		PaymentInfo: entities.PaymentInfo{
			Success:   resp.Success,
			OrderID:   resp.OrderID,
			PaymentID: fmt.Sprint(resp.PaymentID),
		},
		Status:    resp.Status,
		ErrorCode: resp.ErrorCode,
	}
}

func PaymentNotificationFromRest(req apiModels.Notification) entities.PaymentNotification {
	return entities.PaymentNotification{
		Success:     req.Success,
		Amount:      req.Amount,
		ErrorCode:   req.ErrorCode,
		OrderID:     req.OrderID,
		Pan:         req.Pan,
		PaymentID:   req.PaymentID,
		Status:      req.Status,
		TerminalKey: req.TerminalKey,
		Token:       req.Token,
		ExpDate:     req.ExpDate,
		CardID:      req.CardID,
	}
}

func PaymentResponseToLea(e entities.PaymentResponse) rabbitEntities.PaymentResponse {
	return rabbitEntities.PaymentResponse{
		WashID:     e.WashID,
		PostID:     e.PostID,
		OrderID:    e.OrderID,
		UrlPayment: e.UrlPay,
		Failed:     e.Failed,
		Error:      e.Error,
	}
}

func PaymentNotifcationToLea(e entities.PaymentNotificationForLea) rabbitEntities.PaymentNotifcation {
	return rabbitEntities.PaymentNotifcation{
		WashID:  e.WashID,
		PostID:  e.PostID,
		OrderID: e.OrderID,
		Status:  e.Status,
	}
}

func PaymentRequestToSbp(e rabbitEntities.PaymentRequest) entities.PaymentRequest {
	return entities.PaymentRequest{
		WashID:  e.WashID,
		PostID:  e.PostID,
		OrderID: e.OrderID,
		Amount:  e.Amount,
	}
}

func PaymentСancellationRequestToSbp(e rabbitEntities.PaymentСancellationRequest) entities.PaymentСancellationRequest {
	return entities.PaymentСancellationRequest{
		WashID:  e.WashID,
		PostID:  e.PostID,
		OrderID: e.OrderID,
	}
}
