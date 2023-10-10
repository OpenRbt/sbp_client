package repository

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	logicEntities "sbp/internal/logic/entities"
	repConverter "sbp/internal/repository/converter"
	repEntities "sbp/internal/repository/entities"
	"sbp/pkg/bootstrap"

	uuid "github.com/satori/go.uuid"
)

// CreateTransaction ...
func (s *Repository) CreateTransaction(ctx context.Context, transaction logicEntities.TransactionCreate) error {

	_, err := s.db.NewSession(nil).
		InsertInto("transactions").
		Columns("id", "wash_id", "post_id", "amount", "payment_id_bank", "status").
		Record(repEntities.TransactionCreate{
			ID:        transaction.ID,
			WashID:    transaction.WashID,
			PostID:    transaction.PostID,
			Amount:    transaction.Amount,
			PaymentID: transaction.PaymentID,
			Status:    string(transaction.Status),
		}).Exec()

	if err != nil {
		return bootstrap.CustomError(layer, "CreateTransaction", err)
	}
	return nil
}

// UpdateTransaction ...
func (s *Repository) UpdateTransaction(ctx context.Context, transactionUpdate logicEntities.TransactionUpdate) error {
	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return bootstrap.CustomError(layer, "BeginTx", err)
	}

	ID := uuid.NullUUID{UUID: transactionUpdate.ID, Valid: true}

	updateStatement := tx.
		Update("transactions").
		Where("id = ?", ID)

	if transactionUpdate.PaymentID != nil {
		updateStatement = updateStatement.Set("payment_id_bank", transactionUpdate.PaymentID)
	}

	if transactionUpdate.Status != logicEntities.TransactionStatusUnknown {
		updateStatement = updateStatement.Set("status", string(transactionUpdate.Status))
	}

	_, err = updateStatement.ExecContext(ctx)

	if err != nil {
		return bootstrap.CustomError(layer, "UpdateTransaction", err)
	}

	return tx.Commit()
}

// GetTransaction ...
func (s *Repository) GetTransaction(ctx context.Context, orderID uuid.UUID) (logicEntities.Transaction, error) {
	var dbTransaction repEntities.Transaction

	err := s.db.NewSession(nil).
		Select("*").
		From("transactions").
		Where("id = ?", uuid.NullUUID{UUID: orderID, Valid: true}).
		LoadOneContext(ctx, &dbTransaction)

	switch {
	case err == nil:
		return repConverter.ConvertTransactionFromDB(dbTransaction), nil
	default:
		return logicEntities.Transaction{}, bootstrap.CustomError(layer, "GetTransaction", err)
	}
}

// GetTransactionsByStatus ...
func (s *Repository) GetTransactionsByStatus(ctx context.Context, transactionsGet logicEntities.TransactionsGet) ([]logicEntities.Transaction, error) {
	var dbTransactions []repEntities.Transaction

	err := s.db.NewSession(nil).
		Select("*").
		From("transactions").
		Where("status = ?", string(transactionsGet.Status)).
		LoadOneContext(ctx, &dbTransactions)

	switch {
	case err == nil:
		var transactions []logicEntities.Transaction
		for _, t := range dbTransactions {
			transactions = append(transactions, repConverter.ConvertTransactionFromDB(t))
		}
		return transactions, nil
	default:
		return nil, bootstrap.CustomError(layer, "GetTransactionsByStatus", err)
	}
}

// generateNewServiceKey ...
func (s *Repository) generateNewServiceKey() string {
	data := make([]byte, 10)

	_, err := rand.Read(data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", sha256.Sum256(data))
}
