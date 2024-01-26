package rabbitEntities

type MessageType string

const (
	PaymentRequestMessageType             MessageType = "sbp_client_service/payment_request"
	PaymentResponseMessageType            MessageType = "sbp_client_service/payment_response"
	PaymentNotificationMessageType        MessageType = "sbp_client_service/payment_notification"
	PaymentNotificationResponseMessage    MessageType = "sbp_client_service/payment_notification_response"
	PaymentСancellationRequestMessageType MessageType = "sbp_client_service/payment_cancellation_request"
	PaymentErrorMessageType               MessageType = "sbp_client_service/payment_error"

	CreateUserMessageType       MessageType = "admin_service/create_user"
	AdminUserMessageType        MessageType = "admin_service/admin_user"
	OrganizationMessageType     MessageType = "admin_service/organization"
	ServerGroupMessageType      MessageType = "admin_service/server_group"
	RequestAdminDataMessageType MessageType = "admin_service/data"
)
