package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID            uuid.UUID
	WashID        string
	PostID        string
	Amount        int64
	PaymentIDBank string
	Status        TransactionStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TransactionCreate struct {
	ID            uuid.UUID
	WashID        string
	PostID        string
	Amount        int64
	PaymentIDBank string
	Status        TransactionStatus
}

type TransactionsGet struct {
	Status TransactionStatus
}

type TransactionUpdate struct {
	ID            uuid.UUID
	Status        TransactionStatus
	PaymentIDBank *string
}
