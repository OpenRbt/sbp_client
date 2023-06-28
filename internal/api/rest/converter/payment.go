package restconverter

import (
	"fmt"
	logicEntities "sbp/internal/logic/entities"
	openapiEntities "sbp/internal/openapi/models"
)

// СonvertRegisterNotificationFromRest ...
func СonvertRegisterNotificationFromRest(rest openapiEntities.Notification) logicEntities.PaymentRegisterNotification {
	return logicEntities.PaymentRegisterNotification{
		PaymentInfo: logicEntities.PaymentInfo{
			OrderID:   rest.OrderID,
			Success:   rest.Success,
			PaymentID: fmt.Sprintf("%d", rest.PaymentID),
		},
		TerminalKey: rest.TerminalKey,
		Status:      rest.Status,
		ErrorCode:   rest.ErrorCode,
		Amount:      int(rest.Amount),
		CardId:      int(rest.CardID),
		Pan:         rest.Pan,
		ExpDate:     rest.ExpDate,
		Token:       rest.Token,
	}
}
