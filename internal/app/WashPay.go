package app

import uuid "github.com/satori/go.uuid"

type RegisterNotification struct {
	TerminalKey string
	PaymentID   string
	OrderID     string
	Amount      int
	Status      string
	Success     bool
	ErrorCode   string
	ExpDate     string
	Pan         string
	CardId      int
	Token       string
}

type Cancel struct {
	Success   bool
	OrderID   string
	PaymentID string
	Status    string
	ErrorCode string
}

type GetQr struct {
	Success   bool
	OrderID   string
	PaymentID string
	Status    string
	ErrorCode string
	UrlPay    string
	Message   string
}
type Init struct {
	Success   bool
	OrderID   string
	PaymentID string
	Status    string
	Url       string
}

type Transaction struct {
	ID        uuid.UUID
	ServerID  string
	PostID    string
	Amount    int64
	PaymentID string
}

type UpdateTransaction struct {
	ID        uuid.UUID
	Status    *string
	PaymentID *string
}
