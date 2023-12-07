package rabbitEntities

type Message string

const (
	PaymentRequestMessage              Message = "sbp_client_service/payment_request"
	PaymentResponseMessage             Message = "sbp_client_service/payment_response"
	PaymentNotificationMessage         Message = "sbp_client_service/payment_notification"
	PaymentNotificationResponseMessage Message = "sbp_client_service/payment_notification_response"
	PaymentСancellationRequestMessage  Message = "sbp_client_service/payment_cancellation_request"
	PaymentErrorMessage                Message = "sbp_client_service/payment_error"

	AdminUserMessage        Message = "admin_service/admin_user"
	OrganizationMessage     Message = "admin_service/organization"
	ServerGroupMessage      Message = "admin_service/server_group"
	RequestAdminDataMessage Message = "admin_service/data"
)
