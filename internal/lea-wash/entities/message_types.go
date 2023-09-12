package leaentities

import shareBusinessEntities "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"

const (
	// message typess
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
)
