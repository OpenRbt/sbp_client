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
		fallthrough
	case TransactionStatusAuthorized:
		fallthrough
	case TransactionStatusConfirmedNotSynced:
		fallthrough
	case TransactionStatusConfirmed:
		fallthrough
	case TransactionStatusСanceling:
		fallthrough
	case TransactionStatusСanceled:
		fallthrough
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
