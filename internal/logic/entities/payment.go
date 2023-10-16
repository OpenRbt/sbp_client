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

// PaymentNotificationForLea ...
type PaymentNotificationForLea struct {
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
	Status    string
	Url       string
	ErrorCode string
	Message   string
	Details   string
}

// PaymentCancel ...
type PaymentCancel struct {
	PaymentInfo
	Status    string
	ErrorCode string
	Message   string
	Details   string
}

// PaymentGetQr ...
type PaymentGetQr struct {
	PaymentInfo
	UrlPay    string
	ErrorCode string
	Message   string
	Details   string
}

// PaymentNotification ...
type PaymentNotification struct {
	Amount      int64
	ErrorCode   string
	OrderID     string
	Pan         string
	PaymentID   int64
	Status      string
	Success     bool
	TerminalKey string
	Token       string
	CardID      *int64
	ExpDate     *string
}

type InitPaymentResp struct {
	PaymentCreds     PaymentCreds
	TerminalPassword string
}
