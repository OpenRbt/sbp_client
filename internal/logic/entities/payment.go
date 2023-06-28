package entities

// PaymentCreds ...
type PaymentCreds struct {
	TerminalKey string
	PaymentID   string
}

// PaymentCreate ...
type PaymentCreate struct {
	TerminalKey string
	Amount      int64
	OrderID     string
}

// PaymentInfo
type PaymentInfo struct {
	PaymentID string `json:"payment_id"`
	Success   bool   `json:"success"`
	OrderID   string `json:"order_id"`
}

// PaymentInit ...
type PaymentInit struct {
	PaymentInfo
	Status string
	Url    string
}

// PaymentCancel ...
type PaymentCancel struct {
	PaymentInfo
	Status    string
	ErrorCode string
}

// PaymentGetQr ...
type PaymentGetQr struct {
	PaymentInfo
	ErrorCode string
	UrlPay    string
	Message   string
}

// PaymentRegisterNotification ...
type PaymentRegisterNotification struct {
	PaymentInfo
	TerminalKey string `json:"terminal_key"`
	Status      string `json:"status"`
	Amount      int    `json:"amount"`
	CardId      int    `json:"card_id"`
	ErrorCode   string `json:"error_code"`
	ExpDate     string `json:"exp_date"`
	Pan         string `json:"pan"`
	Token       string `json:"token"`
}

type InitPaymentResp struct {
	PaymentCreds     PaymentCreds
	TerminalPassword string
}
