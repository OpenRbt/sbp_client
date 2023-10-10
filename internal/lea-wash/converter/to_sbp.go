package leaentities

import (
	logicEntities "sbp/internal/logic/entities"

	leEntities "sbp/internal/lea-wash/entities"
)

// PaymentRequestToSbp ...
func PaymentRequestToSbp(e leEntities.PaymentRequest) logicEntities.PaymentRequest {
	return logicEntities.PaymentRequest{
		WashID:  e.WashID,
		PostID:  e.PostID,
		OrderID: e.OrderID,
		Amount:  e.Amount,
	}
}

// PaymentСancellationRequestToSbp ...
func PaymentСancellationRequestToSbp(e leEntities.PaymentСancellationRequest) logicEntities.PaymentСancellationRequest {
	return logicEntities.PaymentСancellationRequest{
		WashID:  e.WashID,
		PostID:  e.PostID,
		OrderID: e.OrderID,
	}
}
