package repository

import (
	"context"
	"errors"
	"sbp/internal/conversions"
	"sbp/internal/entities"
	"sbp/internal/helpers"
	"sbp/internal/repository/models"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

func (s *Repository) TransactionsList(ctx context.Context, filter entities.TransactionFilter) (entities.Page[entities.TransactionForPage], error) {
	query := s.db.NewSession(nil).
		Select("count(*)").
		From(dbr.I("transactions").As("t")).
		Join(dbr.I("washes").As("w"), "t.wash_id :: uuid = w.id").
		Join(dbr.I("wash_groups").As("g"), "w.group_id = g.id").
		Join(dbr.I("organizations").As("o"), "g.organization_id = o.id")

	transactionFilterBuild(query, filter)

	var count int64
	err := query.LoadOneContext(ctx, &count)
	if err != nil {
		return entities.Page[entities.TransactionForPage]{}, err
	}

	query = s.db.NewSession(nil).
		Select("t.id, t.post_id, t.amount, t.status, t.created_at, w.id as wash_id, w.title as wash_title, w.deleted as wash_deleted,g.id as group_id, g.name as group_name, g.deleted as group_deleted, o.id as org_id, o.name as org_name, o.deleted as org_deleted").
		From(dbr.I("transactions").As("t")).
		Join(dbr.I("washes").As("w"), "t.wash_id :: uuid = w.id").
		Join(dbr.I("wash_groups").As("g"), "w.group_id = g.id").
		Join(dbr.I("organizations").As("o"), "g.organization_id = o.id").
		OrderDesc("created_at").
		Paginate(uint64(filter.Page()), uint64(filter.PageSize()))

	transactionFilterBuild(query, filter)

	var transactions []models.Transaction
	_, err = query.LoadContext(ctx, &transactions)
	if err != nil {
		return entities.Page[entities.TransactionForPage]{}, err
	}

	return entities.NewPage(conversions.TransactionsFromDB(transactions), filter.Filter, count), nil
}

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

func transactionFilterBuild(query *dbr.SelectStmt, filter entities.TransactionFilter) {
	if filter.PostID != nil {
		query.Where("t.post_id = ?", filter.PostID)
	}
	if filter.WashID != nil {
		query.Where("wash_id = ?", filter.WashID)
	}
	if filter.GroupID != nil {
		query.Where("group_id = ?", filter.GroupID)
	}
	if filter.OrganizationID != nil {
		query.Where("organization_id = ?", filter.OrganizationID)
	}
}
