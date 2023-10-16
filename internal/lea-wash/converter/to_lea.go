package leaentities

import (
	logicEntities "sbp/internal/logic/entities"

	leEntities "sbp/internal/lea-wash/entities"
)

// PaymentResponse ...
func PaymentResponseToLea(e logicEntities.PaymentResponse) leEntities.PaymentResponse {
	return leEntities.PaymentResponse{
		WashID:     e.WashID,
		PostID:     e.PostID,
		OrderID:    e.OrderID,
		UrlPayment: e.UrlPay,
		Failed:     e.Failed,
		Error:      e.Error,
	}
}

// PaymentNotifcation ...
func PaymentNotifcationToLea(e logicEntities.PaymentNotificationForLea) leEntities.PaymentNotifcation {
	return leEntities.PaymentNotifcation{
		WashID:  e.WashID,
		PostID:  e.PostID,
		OrderID: e.OrderID,
		Status:  e.Status,
	}
}
