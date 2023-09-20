package tinkoff

import (
	"context"
	logic "sbp/internal/logic"
	logicEntities "sbp/internal/logic/entities"
	tinkoffClient "sbp/internal/pay-client/tinkoff/client"
	tinkoffOperations "sbp/internal/pay-client/tinkoff/client/operations"
	tinkoffConverter "sbp/internal/pay-client/tinkoff/converter"
	tinkoffEntities "sbp/internal/pay-client/tinkoff/models"
	"time"

	httpTransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"go.uber.org/zap"
)

// check 'PayClient' struct = logic interface 'PayClient'
var _ = logic.PayClient(&PayClient{})

// PayClient ...
type PayClient struct {
	l                          *zap.SugaredLogger
	client                     tinkoffClient.TinkoffAPI
	paymentURLExpirationPeriod time.Duration
}

// NewPayClient ...
func NewPayClient(l *zap.SugaredLogger, paymentURLExpirationPeriod time.Duration) (*PayClient, error) {
	payClient := &PayClient{
		l:                          l,
		paymentURLExpirationPeriod: paymentURLExpirationPeriod,
	}
	transport := httpTransport.New("securepay.tinkoff.ru", "/v2", []string{"https"})
	c := tinkoffClient.New(transport, strfmt.Default)
	payClient.client = *c

	return payClient, nil
}

// Init ...
func (pc *PayClient) Init(req logicEntities.PaymentCreate) (logicEntities.PaymentInit, error) {
	ctx := context.TODO()
	location := time.FixedZone("UTC+03:00", 3*60*60)
	timeInZone := time.Now().Add(pc.paymentURLExpirationPeriod).In(location)
	iso8601Format := "2006-01-02T15:04:05-07:00"
	redirectDueDate := timeInZone.Format(iso8601Format)

	res, err := pc.client.Operations.Init(&tinkoffOperations.InitParams{
		Body: &tinkoffEntities.Init{
			Amount:          &req.Amount,
			OrderID:         &req.OrderID,
			TerminalKey:     &req.TerminalKey,
			RedirectDueDate: redirectDueDate,
		},
		Context:    ctx,
		HTTPClient: nil,
	})
	if err != nil {
		return logicEntities.PaymentInit{}, err
	}
	model := tinkoffConverter.ConvertInitFromResponse(*res.Payload)

	return model, nil
}

// GetQr ...
func (pc *PayClient) GetQr(paymentCreds logicEntities.PaymentCreds, password string) (logicEntities.PaymentGetQr, error) {
	// generate token
	tokkenGenerator, err := NewTokkenGenerator(password)
	if err != nil {
		return logicEntities.PaymentGetQr{}, err
	}

	if err != nil {
		return logicEntities.PaymentGetQr{}, err
	}
	body := tinkoffEntities.GetQr{
		PaymentID:   paymentCreds.PaymentID,
		TerminalKey: paymentCreds.TerminalKey,
		Token:       "",
	}
	token := tokkenGenerator.generateToken(body, "json")
	body.Token = token

	// get QR
	ctx := context.TODO()
	req := tinkoffOperations.GetQrParams{
		Body:       &body,
		Context:    ctx,
		HTTPClient: nil,
	}
	resp, err := pc.client.Operations.GetQr(&req)
	if err != nil {
		return logicEntities.PaymentGetQr{}, err
	}
	model := tinkoffConverter.ConvertGetQrFromResponse(*resp.Payload)

	return model, nil
}

// Cancel ...
func (pc *PayClient) Cancel(req logicEntities.PaymentCreds, password string) (logicEntities.PaymentCancel, error) {
	// generate token
	tokkenGenerator, err := NewTokkenGenerator(password)
	if err != nil {
		return logicEntities.PaymentCancel{}, err
	}
	tinkoffCancel := tinkoffEntities.Cancel{
		PaymentID:   req.PaymentID,
		TerminalKey: req.TerminalKey,
		Token:       "",
	}
	token := tokkenGenerator.generateToken(tinkoffCancel, "json")

	ctx := context.TODO()
	res, err := pc.client.Operations.Cancel(&tinkoffOperations.CancelParams{
		Body: &tinkoffEntities.Cancel{
			PaymentID:   req.PaymentID,
			TerminalKey: req.TerminalKey,
			Token:       token,
		},
		Context:    ctx,
		HTTPClient: nil,
	})
	if err != nil {
		return logicEntities.PaymentCancel{}, err
	}
	model := tinkoffConverter.ConvertCancelFromResponse(*res.Payload)

	return model, nil
}

// IsNotificationCorrect ...
func (pc *PayClient) IsNotificationCorrect(req logicEntities.PaymentRegisterNotification, password string) bool {
	// generate token
	tokkenGenerator, err := NewTokkenGenerator(password)
	if err != nil {
		pc.l.Errorf("IsNotificationCorrect error: %s", err.Error())
		return false
	}

	// paymentRegisterNotification ...
	paymentRegisterNotification := tinkoffConverter.ConvertNotificationFromRequest(req)
	return tokkenGenerator.checkToken("json", paymentRegisterNotification, req.Token)
}
