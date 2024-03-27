package rest

import (
	"net/http"
	"path"

	"sbp/internal/api/handlers"
	"sbp/internal/app"
	"sbp/internal/config"
	entities "sbp/internal/entities"
	"sbp/internal/infrastructure/firebase"
	swaggerRestapi "sbp/openapi/restapi"
	swaggerOperations "sbp/openapi/restapi/operations"
	"sbp/openapi/restapi/operations/notifications"
	"sbp/openapi/restapi/operations/payments"
	"sbp/openapi/restapi/operations/standard"
	"sbp/openapi/restapi/operations/transactions"
	"sbp/openapi/restapi/operations/washes"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"github.com/sebest/xff"
	"go.uber.org/zap"
)

type restApi struct {
	svc        app.Service
	basePath   string
	logger     *zap.SugaredLogger
	swaggerApi *swaggerOperations.WashSbpAPI
	handler    http.Handler

	firebase *firebase.FirebaseClient
}

func (api *restApi) GetSwaggerApi() *swaggerOperations.WashSbpAPI {
	return api.swaggerApi
}

func (api restApi) GetHandler() http.Handler {
	return api.handler
}

func NewApi(cfg config.RestApiConfig) (*restApi, error) {
	err := cfg.CheckRestApiConfig()
	if err != nil {
		return nil, err
	}

	api := &restApi{
		logger: cfg.Logger,
		svc:    cfg.Svc,

		firebase: cfg.Firebase,
	}

	err = api.route()
	if err != nil {
		return nil, err
	}

	err = api.setMiddleware()
	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *restApi) route() error {
	// swagger spec
	swaggerSpec, err := loads.Embedded(swaggerRestapi.SwaggerJSON, swaggerRestapi.FlatSwaggerJSON)
	if err != nil {
		return errors.Wrap(err, "failed to load embedded swagger spec")
	}
	if api.basePath == "" {
		api.basePath = swaggerSpec.BasePath()
	}
	swaggerSpec.Spec().BasePath = api.basePath

	api.swaggerApi = swaggerOperations.NewWashSbpAPI(swaggerSpec)
	api.swaggerApi.Logger = structlog.New(structlog.KeyUnit, "swagger").Printf
	api.swaggerApi.AuthKeyAuth = api.firebase.Auth

	handler := handlers.NewHandler(api.logger, api.svc)

	api.swaggerApi.StandardHealthcheckHandler = standard.HealthcheckHandlerFunc(healthCheck)

	// wash
	api.swaggerApi.WashesCreateWashHandler = washes.CreateWashHandlerFunc(handler.CreateWash)
	api.swaggerApi.WashesUpdateWashHandler = washes.UpdateWashHandlerFunc(handler.UpdateWash)
	api.swaggerApi.WashesDeleteWashHandler = washes.DeleteWashHandlerFunc(handler.DeleteWash)
	api.swaggerApi.WashesGetWashByIDHandler = washes.GetWashByIDHandlerFunc(handler.GetWashByID)
	api.swaggerApi.WashesGetWashesHandler = washes.GetWashesHandlerFunc(handler.GetWashes)

	api.swaggerApi.WashesAssignWashToGroupHandler = washes.AssignWashToGroupHandlerFunc(handler.AssignWashToGroup)

	// payment
	api.swaggerApi.PaymentsInitPaymentHandler = payments.InitPaymentHandlerFunc(handler.InitPayment)
	api.swaggerApi.PaymentsCancelPaymentHandler = payments.CancelPaymentHandlerFunc(handler.CancelPayment)
	api.swaggerApi.NotificationsReceiveNotificationHandler = notifications.ReceiveNotificationHandlerFunc(handler.ReceiveNotification)

	// transactions
	api.swaggerApi.TransactionsGetTransactionsHandler = transactions.GetTransactionsHandlerFunc(handler.GetTransactions)

	return nil
}

func healthCheck(params standard.HealthcheckParams, profile *entities.Auth) standard.HealthcheckResponder {
	return standard.NewHealthcheckOK().WithPayload(&standard.HealthcheckOKBody{Ok: true})
}

func (api *restApi) setMiddleware() error {
	middlewares := func(handler http.Handler) http.Handler {
		return handler
		// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// 	return
		// }

		// fmt.Printf("Request Body: %s\n", string(body))

		// r.Body = io.NopCloser(bytes.NewBuffer(body))

		// handler.ServeHTTP(w, r)
		// })
	}

	handler := api.swaggerApi.Serve(middlewares)

	redocOpts := middleware.RedocOpts{
		BasePath: api.basePath,
		SpecURL:  path.Join(api.basePath, "/swagger.json"),
	}
	docMiddleware := middleware.Redoc(redocOpts, handler)

	truePathMiddleware := middleware.Spec(api.basePath, swaggerRestapi.FlatSwaggerJSON, docMiddleware)

	loggerFunc := makeLogger(api.basePath, api.logger)
	accesslog := makeAccessLog(api.basePath, api.logger)
	loggerMiddleware := accesslog(truePathMiddleware)

	recoveryMiddleware := recovery(loggerMiddleware, api.logger)

	noCacheMiddleware := noCache(recoveryMiddleware)

	lMiddleware := loggerFunc(noCacheMiddleware)

	xffmw, _ := xff.Default()
	finishedMiddleware := xffmw.Handler(lMiddleware)
	// The middleware executes after serving /swagger.json and routing,
	// but before authentication, binding and validation.
	// middlewares := func(handler http.Handler) http.Handler {
	//	safePath := map[string]bool{}
	//	isSafe := func(r *http.Request) bool { return safePath[r.URL.Path] }
	//	//forbidCSRF := makeForbidCSRF(isSafe)
	//
	//	return forbidCSRF(handler)
	//}

	api.handler = finishedMiddleware
	return nil
}
