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

type TransactionForPage struct {
	ID           uuid.UUID
	PostID       int64
	Amount       int64
	Status       TransactionStatus
	CreatedAt    time.Time
	Wash         SimpleWash
	Group        SimpleGroup
	Organization SimpleOrganization
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

type TransactionFilter struct {
	Filter
	OrganizationID *uuid.UUID
	GroupID        *uuid.UUID
	WashID         *uuid.UUID
	PostID         *int64
	Status         *TransactionStatus
}
