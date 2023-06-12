package pay

import (
	"context"
	"fmt"
	"sbp/internal/app"
	"sbp/internal/conversions"
	"sbp/internal/transport/pay/client/operations"
	"sbp/internal/transport/pay/models"
)

func (svc *Service) Init(terminalKey string, amount int64, orderID string) (app.Init, error) {
	fmt.Println("Start INIT")

	ctx := context.TODO()

	res, err := svc.intApi.Operations.Init(&operations.InitParams{
		Body: &models.Init{
			Amount:      &amount,
			OrderID:     &orderID,
			TerminalKey: &terminalKey,
		},
		Context:    ctx,
		HTTPClient: nil,
	})

	if err != nil {
		return app.Init{}, err
	}

	model := conversions.InitFromResponse(*res.Payload)

	return model, nil
}

func (svc *Service) GetQr(terminalKey string, paymentId string, token string) (app.GetQr, error) {

	ctx := context.TODO()

	res, err := svc.intApi.Operations.GetQr(&operations.GetQrParams{
		Body: &models.GetQr{
			PaymentID:   paymentId,
			TerminalKey: terminalKey,
			Token:       token,
		},
		Context:    ctx,
		HTTPClient: nil,
	})

	if err != nil {
		return app.GetQr{}, err
	}
	fmt.Println("HTTP ", res.Payload.ErrorCode)

	model := conversions.GetQrFromResponse(*res.Payload)

	return model, nil
}

func (svc *Service) Cancel(terminalKey string, paymentId string, token string) (app.Cancel, error) {

	ctx := context.TODO()

	res, err := svc.intApi.Operations.Cancel(&operations.CancelParams{
		Body: &models.Cancel{
			PaymentID:   paymentId,
			TerminalKey: terminalKey,
			Token:       token,
		},
		Context:    ctx,
		HTTPClient: nil,
	})

	if err != nil {
		return app.Cancel{}, err
	}

	model := conversions.CancelFromResponse(*res.Payload)

	return model, nil
}
