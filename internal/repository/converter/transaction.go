package repconverter

import (
	logicEntities "sbp/internal/logic/entities"
	repEntities "sbp/internal/repository/entities"
	"time"
)

func ConvertTransactionFromDB(dbTransaction repEntities.Transaction) logicEntities.Transaction {
	return logicEntities.Transaction{
		ID:         dbTransaction.ID.UUID,
		ServerID:   dbTransaction.ServerID,
		PostID:     dbTransaction.PostID,
		Amount:     dbTransaction.Amount,
		PaymentID:  dbTransaction.PaymentID,
		Status:     dbTransaction.Status,
		DataCreate: time.Time{},
	}
}
