package rest

import (
	"net/http"
	"path"

	logicEntities "sbp/internal/logic/entities"
	swaggerRestapi "sbp/openapi/restapi"
	swaggerOperations "sbp/openapi/restapi/operations"
	"sbp/openapi/restapi/operations/standard"
	"sbp/openapi/restapi/operations/wash_servers"

	restHandlers "sbp/internal/api/rest/handlers"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"github.com/sebest/xff"
	"go.uber.org/zap"
)

// RestApiConfig ...
type RestApiConfig struct {
	Logger *zap.SugaredLogger
	Logic  restHandlers.Logic
	// TODO: extend with services
}

// checkRestApiConfig ...
func checkRestApiConfig(conf RestApiConfig) error {
	if conf.Logger == nil {
		return errors.New("api logger is empty")
	}
	if conf.Logic == nil {
		return errors.New("api logic is empty")
	}

	return nil
}

// restApi ...
type restApi struct {
	logic      restHandlers.Logic
	basePath   string
	logger     *zap.SugaredLogger
	swaggerApi *swaggerOperations.WashSbpAPI
	handler    http.Handler
}

// GetSwaggerApi ...
func (api *restApi) GetSwaggerApi() *swaggerOperations.WashSbpAPI {
	return api.swaggerApi
}

// GetHandler ...
func (api restApi) GetHandler() http.Handler {
	return api.handler
}

// NewApi ...
func NewApi(restApiConfig RestApiConfig) (*restApi, error) {
	err := checkRestApiConfig(restApiConfig)
	if err != nil {
		return nil, err
	}
	// init api
	api := &restApi{
		logger: restApiConfig.Logger,
		logic:  restApiConfig.Logic,
	}

	// route
	err = api.route()
	if err != nil {
		return nil, err
	}

	// middleware
	err = api.setMiddleware()
	if err != nil {
		return nil, err
	}

	return api, nil
}

// route ...
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

	// init swagger
	api.swaggerApi = swaggerOperations.NewWashSbpAPI(swaggerSpec)
	api.swaggerApi.Logger = structlog.New(structlog.KeyUnit, "swagger").Printf

	// handlers
	handler := restHandlers.NewHandler(api.logger, api.logic)

	// health check
	api.swaggerApi.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(healthCheck)

	// handlers mapping
	// auth
	api.swaggerApi.AuthKeyAuth = handler.Auth
	// api.swaggerApi.WashServersSignupHandler = wash_servers.SignupHandlerFunc(handler.SignUP)

	// wash server
	api.swaggerApi.WashServersCreateHandler = wash_servers.CreateHandlerFunc(handler.CreateWashServer)
	api.swaggerApi.WashServersUpdateHandler = wash_servers.UpdateHandlerFunc(handler.UpdateWashServer)
	api.swaggerApi.WashServersDeleteHandler = wash_servers.DeleteHandlerFunc(handler.DeleteWashServer)
	api.swaggerApi.WashServersGetWashServerHandler = wash_servers.GetWashServerHandlerFunc(handler.GetWashServer)

	// payment
	// api.swaggerApi.WashServersPayHandler = wash_servers.PayHandlerFunc(handler.Pay)
	// api.swaggerApi.WashServersCancelHandler = wash_servers.CancelHandlerFunc(handler.Cancel)
	api.swaggerApi.WashServersNotificationHandler = wash_servers.NotificationHandlerFunc(handler.Notif)
	return nil
}

// healthCheck ...
func healthCheck(params standard.HealthCheckParams, profile *logicEntities.Auth) standard.HealthCheckResponder {
	return standard.NewHealthCheckOK().WithPayload(&standard.HealthCheckOKBody{Ok: true})
}

// setMiddleware ...
func (api *restApi) setMiddleware() error {
	//
	middlewares := func(handler http.Handler) http.Handler {
		return handler
	}

	handler := api.swaggerApi.Serve(middlewares)

	// doc middleware
	redocOpts := middleware.RedocOpts{
		BasePath: api.basePath,
		SpecURL:  path.Join(api.basePath, "/swagger.json"),
	}
	docMiddleware := middleware.Redoc(redocOpts, handler)

	// true path middleware
	truePathMiddleware := middleware.Spec(api.basePath, swaggerRestapi.FlatSwaggerJSON, docMiddleware)

	// logger middleware
	loggerFunc := makeLogger(api.basePath, api.logger)
	accesslog := makeAccessLog(api.basePath, api.logger)
	loggerMiddleware := accesslog(truePathMiddleware)

	// recovery middleware
	recoveryMiddleware := recovery(loggerMiddleware, api.logger)

	// no cache middleware
	noCacheMiddleware := noCache(recoveryMiddleware)

	// logger
	lMiddleware := loggerFunc(noCacheMiddleware)

	// forwarded HTTP Extension
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
