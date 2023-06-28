package entities

import (
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type transactionStatus string

const (
	TransactionStatusNew                transactionStatus = "new"
	TransactionStatusAuthorized         transactionStatus = "authorized"
	TransactionStatusConfirmedNotSynced transactionStatus = "confirmed_not_synced"
	TransactionStatusConfirmed          transactionStatus = "confirmed"
	TransactionStatusСanceling          transactionStatus = "canceling"
	TransactionStatusСanceled           transactionStatus = "canceled"
	TransactionStatusUnknown            transactionStatus = "unknown"
)

var knownStatuses []transactionStatus = []transactionStatus{
	TransactionStatusNew,
	TransactionStatusAuthorized,
	TransactionStatusConfirmedNotSynced,
	TransactionStatusConfirmed,
	TransactionStatusСanceling,
	TransactionStatusСanceled,
}

// TransactionStatusFromString ...
func TransactionStatusFromString(s string) (transactionStatus, error) {
	s = strings.ToLower(s)
	for i, k := range knownStatuses {
		if string(k) == s {
			return knownStatuses[i], nil
		}
	}
	return TransactionStatusUnknown, fmt.Errorf("transaction status is unknown: %s", s)
}

// TransactionCreate ...
type TransactionCreate struct {
	ID        uuid.UUID
	ServerID  string
	PostID    string
	Amount    int64
	PaymentID string
	Status    transactionStatus
}

// TransactionsGet ...
type TransactionsGet struct {
	Status transactionStatus
}

// Transaction ...
type Transaction struct {
	ID         uuid.UUID
	ServerID   string
	PostID     string
	Amount     int64
	PaymentID  string
	Status     string
	DataCreate time.Time
}

// TransactionUpdate ...
type TransactionUpdate struct {
	ID        uuid.UUID
	Status    transactionStatus
	PaymentID *string
}
