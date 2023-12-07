package app

import (
	"sbp/internal/entities"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

type (
	PaymentRepository interface {
		GetTransactionByOrderID(ctx Ctx, id uuid.UUID) (entities.Transaction, error)
		CreateTransaction(ctx context.Context, createTransaction entities.TransactionCreate) error
		UpdateTransaction(ctx context.Context, updateTransaction entities.TransactionUpdate) error

		GetTransactionsByStatus(ctx context.Context, transactionsGet entities.TransactionsGet) ([]entities.Transaction, error)
	}

	PaymentService interface {
		InitPayment(ctx context.Context, req entities.PaymentRequest) (*entities.PaymentResponse, error)
		CancelPayment(ctx context.Context, req entities.PaymentСancellationRequest) (resendNeeded bool, err error)
		ReceiveNotification(ctx context.Context, notification entities.PaymentNotification) error

		SyncAllPayments(ctx context.Context) error
	}

	PaymentClient interface {
		InitPayment(req entities.PaymentCreate) (entities.PaymentInit, error)
		GetQr(req entities.PaymentCreds, password string) (entities.PaymentGetQr, error)
		CancelPayment(req entities.PaymentCreds, password string) (entities.PaymentCancel, error)
	}
)
