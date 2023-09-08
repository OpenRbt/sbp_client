package repconverter

import (
	logicEntities "sbp/internal/logic/entities"
	repEntities "sbp/internal/repository/entities"
)

func ConvertTransactionFromDB(dbTransaction repEntities.Transaction) logicEntities.Transaction {
	return logicEntities.Transaction{
		ID:        dbTransaction.ID,
		WashID:    dbTransaction.WashID,
		PostID:    dbTransaction.PostID,
		Amount:    dbTransaction.Amount,
		PaymentID: dbTransaction.PaymentID,
		Status:    logicEntities.TransactionStatusFromString(dbTransaction.Status),
		CreatedAt: dbTransaction.CreatedAt,
		UpdatedAt: dbTransaction.UpdatedAt,
	}
}
