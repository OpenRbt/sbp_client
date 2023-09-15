package entities

import "strings"

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

// validTransactionStatuses ...
var validTransactionStatuses map[transactionStatus]bool = map[transactionStatus]bool{
	TransactionStatusNew:                true,
	TransactionStatusAuthorized:         true,
	TransactionStatusConfirmedNotSynced: true,
	TransactionStatusConfirmed:          true,
	TransactionStatusСanceling:          true,
	TransactionStatusСanceled:           true,
}

// ValidateTransactionStatus ...
func ValidateTransactionStatus(r transactionStatus) bool {
	return validTransactionStatuses[r]
}

// String ...
func (s transactionStatus) String() string {
	return string(s)
}

// TransactionStatusFromString ...
func TransactionStatusFromString(s string) transactionStatus {
	s = strings.ToLower(s)
	status := transactionStatus(s)
	ok := validTransactionStatuses[status]
	if !ok {
		status = TransactionStatusUnknown
	}
	return status
}
