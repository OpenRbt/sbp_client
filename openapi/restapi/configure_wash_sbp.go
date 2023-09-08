// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"sbp/internal/logic/entities"
	"sbp/openapi/restapi/operations"
	"sbp/openapi/restapi/operations/standard"
	"sbp/openapi/restapi/operations/wash"
)


func configureFlags(api *operations.WashSbpAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WashSbpAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.AuthKeyAuth == nil {
		api.AuthKeyAuth = func(token string) (*entities.AuthExtended, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.WashCancelHandler == nil {
		api.WashCancelHandler = wash.CancelHandlerFunc(func(params wash.CancelParams, principal *entities.AuthExtended) wash.CancelResponder {
			return wash.CancelNotImplemented()
		})
	}
	if api.WashCreateHandler == nil {
		api.WashCreateHandler = wash.CreateHandlerFunc(func(params wash.CreateParams, principal *entities.AuthExtended) wash.CreateResponder {
			return wash.CreateNotImplemented()
		})
	}
	if api.WashDeleteHandler == nil {
		api.WashDeleteHandler = wash.DeleteHandlerFunc(func(params wash.DeleteParams, principal *entities.AuthExtended) wash.DeleteResponder {
			return wash.DeleteNotImplemented()
		})
	}
	if api.WashGetWashHandler == nil {
		api.WashGetWashHandler = wash.GetWashHandlerFunc(func(params wash.GetWashParams, principal *entities.AuthExtended) wash.GetWashResponder {
			return wash.GetWashNotImplemented()
		})
	}
	if api.StandardHealthCheckHandler == nil {
		api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal *entities.AuthExtended) standard.HealthCheckResponder {
			return standard.HealthCheckNotImplemented()
		})
	}
	if api.WashListHandler == nil {
		api.WashListHandler = wash.ListHandlerFunc(func(params wash.ListParams, principal *entities.AuthExtended) wash.ListResponder {
			return wash.ListNotImplemented()
		})
	}
	if api.WashNotificationHandler == nil {
		api.WashNotificationHandler = wash.NotificationHandlerFunc(func(params wash.NotificationParams, principal *entities.AuthExtended) wash.NotificationResponder {
			return wash.NotificationNotImplemented()
		})
	}
	if api.WashPayHandler == nil {
		api.WashPayHandler = wash.PayHandlerFunc(func(params wash.PayParams, principal *entities.AuthExtended) wash.PayResponder {
			return wash.PayNotImplemented()
		})
	}
	if api.WashSignupHandler == nil {
		api.WashSignupHandler = wash.SignupHandlerFunc(func(params wash.SignupParams, principal *entities.AuthExtended) wash.SignupResponder {
			return wash.SignupNotImplemented()
		})
	}
	if api.WashUpdateHandler == nil {
		api.WashUpdateHandler = wash.UpdateHandlerFunc(func(params wash.UpdateParams, principal *entities.AuthExtended) wash.UpdateResponder {
			return wash.UpdateNotImplemented()
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
