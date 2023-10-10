package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Transaction ...
type Transaction struct {
	ID        uuid.UUID
	WashID    string
	PostID    string
	Amount    int64
	PaymentID string
	Status    transactionStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TransactionCreate ...
type TransactionCreate struct {
	ID        uuid.UUID
	WashID    string
	PostID    string
	Amount    int64
	PaymentID string
	Status    transactionStatus
}

// TransactionsGet ...
type TransactionsGet struct {
	Status transactionStatus
}

// TransactionUpdate ...
type TransactionUpdate struct {
	ID        uuid.UUID
	Status    transactionStatus
	PaymentID *string
}
