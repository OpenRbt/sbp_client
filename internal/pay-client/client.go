package tinkoff

import (
	"context"
	"fmt"
	"reflect"
	logic "sbp/internal/logic"
	logicEntities "sbp/internal/logic/entities"
	converter "sbp/internal/pay-client/converter"
	tinkoffClient "sbp/internal/tinkoff/client"
	tinkoffOperations "sbp/internal/tinkoff/client/operations"
	tinkoffEntities "sbp/internal/tinkoff/models"
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
	model := converter.PaymentInitResponseToLogic(*res.Payload)

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

	// token
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
	model := converter.GetQrResponseToLogic(*resp.Payload)

	return model, nil
}

func structToMap(input interface{}) (map[string]string, error) {
	result := make(map[string]string)
	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Input must be a struct or a pointer to a struct")
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fieldValue := fmt.Sprintf("%v", field.Interface())
		result[fieldName] = fieldValue
	}

	return result, nil
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
	model := converter.PaymentCancelResponseToLogic(*res.Payload)

	return model, nil
}

// IsNotificationCorrect ...
func (pc *PayClient) IsNotificationCorrect(notification logicEntities.PaymentNotification, password string) bool {
	// generate token
	tokkenGenerator, err := NewTokkenGenerator(password)
	if err != nil {
		pc.l.Errorf("IsNotificationCorrect error: %s", err.Error())
		return false
	}

	notificationWithJsonTag := converter.PaymentNotificationToPayClient(notification)
	// paymentRegisterNotification ...
	isTokenValid := tokkenGenerator.checkToken("json", notificationWithJsonTag, notification.Token)
	return isTokenValid
}
