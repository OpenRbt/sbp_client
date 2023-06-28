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
	MessageTypePaymentRequest  shareBusinessEntities.MessageType = "sbp_client_service/payment_request"
	MessageTypePaymentResponse shareBusinessEntities.MessageType = "sbp_client_service/payment_response"

	MessageTypePaymentNotification         shareBusinessEntities.MessageType = "sbp_client_service/payment_notification"
	MessageTypePaymentNotificationResponse shareBusinessEntities.MessageType = "sbp_client_service/payment_notification_response"

	MessageTypePaymentСancellationRequest shareBusinessEntities.MessageType = "sbp_client_service/payment_cancellation_request"

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
	ServerID  string `json:"server_id"`
	PostID    string `json:"post_id"`
	OrderID   string `json:"order_id"`
	ErrorCode int64  `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
}

// PayRequest ...
type PayRequest struct {
	Amount     int64  `json:"amount"`
	ServerID   string `json:"server_id"`
	PostID     string `json:"post_id"`
	OrderID    string `json:"order_id,omitempty"`
	ServiceKey string `json:"service_key"`
}

// PayResponse ...
type PayResponse struct {
	PostID  string `json:"post_id"`
	OrderID string `json:"order_id"`
	UrlPay  string `json:"url_pay"`
}

// PayСancellationRequest ...
type PayСancellationRequest struct {
	ServerID   string `json:"server_id"`
	PostID     string `json:"post_id"`
	ServiceKey string `json:"service_key"`
	OrderID    string `json:"order_id"`
}

// PayNotifcation ...
type PayNotifcation struct {
	ServerID string `json:"server_id"`
	PostID   string `json:"post_id"`
	OrderID  string `json:"order_id"`
	Status   string `json:"status"`
}
