package server

import (
	"net/http"
	swaggerOperations "sbp/internal/openapi/restapi/operations"
)

// Api ...
type Api interface {
	GetSwaggerApi() *swaggerOperations.WashSbpAPI
	GetHandler() http.Handler
}
