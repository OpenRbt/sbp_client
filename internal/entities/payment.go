package entities

type PaymentRequest struct {
	Amount  int64
	WashID  string
	PostID  string
	OrderID string
}

type PaymentResponse struct {
	WashID  string
	PostID  string
	OrderID string
	UrlPay  string
	Failed  bool
	Error   string
}

type PaymentСancellationRequest struct {
	WashID  string
	PostID  string
	OrderID string
}

type PaymentNotificationForLea struct {
	WashID  string
	PostID  string
	OrderID string
	Status  string
}

type PaymentCreds struct {
	TerminalKey string
	PaymentID   string
}

type PaymentCreate struct {
	TerminalKey string
	Amount      int64
	OrderID     string
}

type PaymentInfo struct {
	PaymentID string
	Success   bool
	OrderID   string
}

type PaymentInit struct {
	PaymentInfo
	Status    string
	Url       string
	ErrorCode string
	Message   string
	Details   string
}

type PaymentCancel struct {
	PaymentInfo
	Status    string
	ErrorCode string
	Message   string
	Details   string
}

type PaymentGetQr struct {
	PaymentInfo
	UrlPay    string
	ErrorCode string
	Message   string
	Details   string
}

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
