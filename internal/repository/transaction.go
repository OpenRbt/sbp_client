package repository

import (
	"context"
	"errors"
	"sbp/internal/entities"
	"sbp/internal/helpers"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

func (s *Repository) CreateTransaction(ctx context.Context, transaction entities.TransactionCreate) error {
	_, err := s.db.NewSession(nil).
		InsertInto("transactions").
		Columns("id", "wash_id", "post_id", "amount", "payment_id_bank", "status").
		Record(transaction).
		Exec()

	if err != nil {
		return helpers.CustomError(layer, "CreateTransaction", err)
	}
	return nil
}

func (s *Repository) UpdateTransaction(ctx context.Context, transactionUpdate entities.TransactionUpdate) error {
	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return helpers.CustomError(layer, "BeginTx", err)
	}

	ID := uuid.NullUUID{UUID: transactionUpdate.ID, Valid: true}

	updateStatement := tx.
		Update("transactions").
		Where("id = ?", ID)

	if transactionUpdate.PaymentIDBank != nil {
		updateStatement = updateStatement.Set("payment_id_bank", transactionUpdate.PaymentIDBank)
	}

	if transactionUpdate.Status != entities.TransactionStatusUnknown {
		updateStatement = updateStatement.Set("status", string(transactionUpdate.Status))
	}

	_, err = updateStatement.ExecContext(ctx)

	if err != nil {
		return helpers.CustomError(layer, "UpdateTransaction", err)
	}

	return tx.Commit()
}

func (s *Repository) GetTransactionByOrderID(ctx context.Context, orderID uuid.UUID) (entities.Transaction, error) {
	var transaction entities.Transaction

	err := s.db.NewSession(nil).
		Select("*").
		From("transactions").
		Where("id = ?", orderID).
		LoadOneContext(ctx, &transaction)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = entities.ErrNotFound
		}

		return entities.Transaction{}, helpers.CustomError(layer, "GetTransactionByOrderID", err)
	}

	return transaction, nil
}

func (s *Repository) GetTransactionsByStatus(ctx context.Context, transactionsGet entities.TransactionsGet) ([]entities.Transaction, error) {
	var transactions []entities.Transaction

	err := s.db.NewSession(nil).
		Select("*").
		From("transactions").
		Where("status = ?", string(transactionsGet.Status)).
		LoadOneContext(ctx, &transactions)

	if err != nil {
		return nil, helpers.CustomError(layer, "GetTransactionsByStatus", err)
	}

	return transactions, nil
}
