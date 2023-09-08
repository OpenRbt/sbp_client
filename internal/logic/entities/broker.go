package entities

import (
	shareBusinessEntities "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
)

const (
	// services
	ServiceLeaCentralWash shareBusinessEntities.Service = "lea_central_wash_service"
	ServiceSbpClient      shareBusinessEntities.Service = "sbp_client_service"

	// routing keys
	RoutingKeySbpClient shareBusinessEntities.RoutingKey = "sbp_client"

	// message types
	// - payment
	MessageTypePaymentRequest  shareBusinessEntities.MessageType = "sbp_client_service/payment_request"
	MessageTypePaymentResponse shareBusinessEntities.MessageType = "sbp_client_service/payment_response"
	// - notification
	MessageTypePaymentNotification         shareBusinessEntities.MessageType = "sbp_client_service/payment_notification"
	MessageTypePaymentNotificationResponse shareBusinessEntities.MessageType = "sbp_client_service/payment_notification_response"
	// - cancellation
	MessageTypePaymentСancellationRequest shareBusinessEntities.MessageType = "sbp_client_service/payment_cancellation_request"
	// - error
	MessageTypePaymentError shareBusinessEntities.MessageType = "sbp_client_service/payment_error"

	// errors
	ErrorPaymentRequestFailed      int64 = iota
	ErrorPaymentСancellationFailed int64 = iota + 1
)

// BrokerMessage ...
type BrokerMessage struct {
	Message     []byte
	Service     shareBusinessEntities.Service
	RoutingKey  shareBusinessEntities.RoutingKey
	MessageType shareBusinessEntities.MessageType
}

// PayError ...
type PayError struct {
	WashID    string `json:"wash_id"`
	PostID    string `json:"post_id"`
	OrderID   string `json:"order_id"`
	ErrorCode int64  `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
}

// PayRequest ...
type PayRequest struct {
	Amount  int64  `json:"amount"`
	WashID  string `json:"wash_id"`
	PostID  string `json:"post_id"`
	OrderID string `json:"order_id,omitempty"`
}

// PayResponse ...
type PayResponse struct {
	WashID  string `json:"wash_id"`
	PostID  string `json:"post_id"`
	OrderID string `json:"order_id"`
	UrlPay  string `json:"url_pay"`
}

// PayСancellationRequest ...
type PayСancellationRequest struct {
	WashID  string `json:"wash_id"`
	PostID  string `json:"post_id"`
	OrderID string `json:"order_id"`
}

// PayNotifcation ...
type PayNotifcation struct {
	WashID  string `json:"wash_id"`
	PostID  string `json:"post_id"`
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}
