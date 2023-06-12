package dal

import (
	"context"
	"sbp/internal/app"
	"sbp/internal/conversions"
	"sbp/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

func (s *Storage) NewTransaction(ctx context.Context, serverID string, postID string, amount int) (app.Transaction, error) {
	var registredTransaction dbmodels.Transaction

	err := s.db.NewSession(nil).
		InsertInto("transactions").
		Columns("server_id", "post_id", "amount", "payment_id_bank").
		Record(dbmodels.AddTransaction{
			ServerID:  serverID,
			PostID:    postID,
			Amount:    amount,
			PaymentID: "",
		}).Returning("id", "server_id", "post_id", "amount", "payment_id_bank").
		LoadContext(ctx, &registredTransaction)

	if err != nil {
		return app.Transaction{}, err
	}
	return conversions.TransactionFromDB(registredTransaction), nil
}

func (s *Storage) UpdateTransaction(ctx context.Context, id uuid.UUID, paymentID *string, status *string) error {
	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	ID := uuid.NullUUID{UUID: id, Valid: true}

	updateStatement := tx.
		Update("transactions").
		Where("id = ?", ID)

	if paymentID != nil {
		updateStatement = updateStatement.Set("payment_id_bank", paymentID)
	}

	if status != nil {
		updateStatement = updateStatement.Set("status", status)
	}

	_, err = updateStatement.ExecContext(ctx)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Storage) GetTransaction(ctx context.Context, orderID uuid.UUID) (app.Transaction, error) {
	var dbTransaction dbmodels.Transaction

	err := s.db.NewSession(nil).
		Select("*").
		From("transactions").
		Where("id = ?", uuid.NullUUID{UUID: orderID, Valid: true}).
		LoadOneContext(ctx, &dbTransaction)

	switch {
	case err == nil:
		return conversions.TransactionFromDB(dbTransaction), err
	default:
		return app.Transaction{}, err
	}
}
