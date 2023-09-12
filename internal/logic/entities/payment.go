package entities

// PaymentRequest ...
type PaymentRequest struct {
	Amount  int64
	WashID  string
	PostID  string
	OrderID string
}

// PaymentResponse ...
type PaymentResponse struct {
	WashID  string
	PostID  string
	OrderID string
	UrlPay  string
	Failed  bool
}

// PaymentСancellationRequest ...
type PaymentСancellationRequest struct {
	WashID  string
	PostID  string
	OrderID string
}

// PaymentNotifcation ...
type PaymentNotifcation struct {
	WashID  string
	PostID  string
	OrderID string
	Status  string
}

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
	PaymentID string
	Success   bool
	OrderID   string
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
	TerminalKey string
	Status      string
	Amount      int
	CardId      int
	ErrorCode   string
	ExpDate     string
	Pan         string
	Token       string
}

type InitPaymentResp struct {
	PaymentCreds     PaymentCreds
	TerminalPassword string
}
