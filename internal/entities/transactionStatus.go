package entities

import "strings"

type TransactionStatus string

const (
	TransactionStatusNew                TransactionStatus = "new"
	TransactionStatusAuthorized         TransactionStatus = "authorized"
	TransactionStatusConfirmedNotSynced TransactionStatus = "confirmed_not_synced"
	TransactionStatusConfirmed          TransactionStatus = "confirmed"
	TransactionStatusСanceling          TransactionStatus = "canceling"
	TransactionStatusСanceled           TransactionStatus = "canceled"
	TransactionStatusRefunded           TransactionStatus = "refunded"
	TransactionStatusUnknown            TransactionStatus = "unknown"
)

func ValidateTransactionStatus(r TransactionStatus) bool {
	switch r {
	case TransactionStatusNew:
	case TransactionStatusAuthorized:
	case TransactionStatusConfirmedNotSynced:
	case TransactionStatusConfirmed:
	case TransactionStatusСanceling:
	case TransactionStatusСanceled:
	case TransactionStatusRefunded:
		return true
	}

	return false
}

func (status TransactionStatus) String() string {
	return string(status)
}

func TransactionStatusFromString(s string) TransactionStatus {
	s = strings.ToLower(s)
	status := TransactionStatus(s)
	if ok := ValidateTransactionStatus(status); !ok {
		status = TransactionStatusUnknown
	}

	return status
}
