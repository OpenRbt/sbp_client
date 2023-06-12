package pay

import (
	tinfoffClient "sbp/internal/transport/pay/client"

	httptransport "github.com/go-openapi/runtime/client"
	"go.uber.org/zap"

	"github.com/go-openapi/strfmt"
)

type Service struct {
	l *zap.SugaredLogger

	intApi tinfoffClient.TinkoffAPI
}

func New(l *zap.SugaredLogger) (svc *Service, err error) {
	svc = &Service{
		l: l,
	}

	initClient := tinfoffClient.New(httptransport.New("securepay.tinkoff.ru", "/v2", []string{"https"}), strfmt.Default)

	svc.intApi = *initClient
	return svc, nil
}
