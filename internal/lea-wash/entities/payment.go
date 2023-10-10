package leaentities

// PaymentRequest ...
type PaymentRequest struct {
	Amount  int64  `json:"amount"`
	WashID  string `json:"wash_id"`
	PostID  string `json:"post_id"`
	OrderID string `json:"order_id"`
}

// PaymentResponse ...
type PaymentResponse struct {
	WashID     string `json:"wash_id"`
	PostID     string `json:"post_id"`
	OrderID    string `json:"order_id"`
	UrlPayment string `json:"url_pay"`
	Failed     bool   `json:"failed"`
}

// PaymentСancellationRequest ...
type PaymentСancellationRequest struct {
	WashID  string `json:"wash_id"`
	PostID  string `json:"post_id"`
	OrderID string `json:"order_id"`
}

// PaymentNotifcation ...
type PaymentNotifcation struct {
	WashID  string `json:"wash_id"`
	PostID  string `json:"post_id"`
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}
