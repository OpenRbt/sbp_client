package conversions

import (
	"sbp/internal/app"
	"sbp/internal/dal/dbmodels"
	"sbp/openapi/models"
)

func TransactionFromDB(dbTransaction dbmodels.Transaction) app.Transaction {
	return app.Transaction{
		ID:        dbTransaction.ID.UUID,
		ServerID:  dbTransaction.ServerID,
		PostID:    dbTransaction.PostID,
		Amount:    dbTransaction.Amount,
		PaymentID: dbTransaction.PaymentID,
	}
}

func RegisterNotificationFromRest(rest models.Notification) app.RegisterNotification {
	return app.RegisterNotification{
		TerminalKey: rest.TerminalKey,
		OrderID:     rest.OrderID,
		Success:     rest.Success,
		Status:      rest.Status,
		PaymentID:   rest.PaymentID,
		ErrorCode:   rest.ErrorCode,
		Amount:      int(rest.Amount),
		CardId:      int(rest.CardID),
		Pan:         rest.Pan,
		ExpDate:     rest.ExpDate,
		Token:       rest.Token,
	}
}
