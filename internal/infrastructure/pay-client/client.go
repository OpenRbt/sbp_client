package tinkoff

import (
	"context"
	"fmt"
	"reflect"
	"sbp/internal/app"
	converter "sbp/internal/conversions"
	"sbp/internal/entities"
	tinkoffClient "sbp/tinkoffapi/client"
	tinkoffOperations "sbp/tinkoffapi/client/operations"
	tinkoffEntities "sbp/tinkoffapi/models"
	"time"

	httpTransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"go.uber.org/zap"
)

var _ = app.PaymentClient(&PayClient{})

type PayClient struct {
	l                          *zap.SugaredLogger
	client                     tinkoffClient.TinkoffAPI
	paymentURLExpirationPeriod time.Duration
}

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

func (svc *PayClient) InitPayment(req entities.PaymentCreate) (entities.PaymentInit, error) {
	ctx := context.TODO()
	location := time.FixedZone("UTC+03:00", 3*60*60)
	timeInZone := time.Now().Add(svc.paymentURLExpirationPeriod).In(location)
	iso8601Format := "2006-01-02T15:04:05-07:00"
	redirectDueDate := timeInZone.Format(iso8601Format)

	res, err := svc.client.Operations.Init(&tinkoffOperations.InitParams{
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
		return entities.PaymentInit{}, err
	}
	model := converter.InitPaymentFromRest(*res.Payload)

	return model, nil
}

func (svc *PayClient) GetQr(paymentCreds entities.PaymentCreds, password string) (entities.PaymentGetQr, error) {
	// generate token
	tokkenGenerator, err := newTokkenGenerator(password)
	if err != nil {
		return entities.PaymentGetQr{}, err
	}

	if err != nil {
		return entities.PaymentGetQr{}, err
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
	resp, err := svc.client.Operations.GetQr(&req)
	if err != nil {
		return entities.PaymentGetQr{}, err
	}
	model := converter.GetQRFromRest(*resp.Payload)

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

func (svc *PayClient) CancelPayment(req entities.PaymentCreds, password string) (entities.PaymentCancel, error) {
	// generate token
	tokkenGenerator, err := newTokkenGenerator(password)
	if err != nil {
		return entities.PaymentCancel{}, err
	}
	tinkoffCancel := tinkoffEntities.Cancel{
		PaymentID:   req.PaymentID,
		TerminalKey: req.TerminalKey,
		Token:       "",
	}

	token := tokkenGenerator.generateToken(tinkoffCancel, "json")

	ctx := context.TODO()
	res, err := svc.client.Operations.Cancel(&tinkoffOperations.CancelParams{
		Body: &tinkoffEntities.Cancel{
			PaymentID:   req.PaymentID,
			TerminalKey: req.TerminalKey,
			Token:       token,
		},
		Context:    ctx,
		HTTPClient: nil,
	})
	if err != nil {
		return entities.PaymentCancel{}, err
	}
	model := converter.CancelPaymentFromRest(*res.Payload)

	return model, nil
}
