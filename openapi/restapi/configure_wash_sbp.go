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
	"sbp/openapi/restapi/operations/wash_servers"
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
		api.AuthKeyAuth = func(token string) (*entities.Auth, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.WashServersCreateHandler == nil {
		api.WashServersCreateHandler = wash_servers.CreateHandlerFunc(func(params wash_servers.CreateParams, principal *entities.Auth) wash_servers.CreateResponder {
			return wash_servers.CreateNotImplemented()
		})
	}
	if api.WashServersDeleteHandler == nil {
		api.WashServersDeleteHandler = wash_servers.DeleteHandlerFunc(func(params wash_servers.DeleteParams, principal *entities.Auth) wash_servers.DeleteResponder {
			return wash_servers.DeleteNotImplemented()
		})
	}
	if api.WashServersGetWashServerHandler == nil {
		api.WashServersGetWashServerHandler = wash_servers.GetWashServerHandlerFunc(func(params wash_servers.GetWashServerParams, principal *entities.Auth) wash_servers.GetWashServerResponder {
			return wash_servers.GetWashServerNotImplemented()
		})
	}
	if api.StandardHealthCheckHandler == nil {
		api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal *entities.Auth) standard.HealthCheckResponder {
			return standard.HealthCheckNotImplemented()
		})
	}
	if api.WashServersListHandler == nil {
		api.WashServersListHandler = wash_servers.ListHandlerFunc(func(params wash_servers.ListParams, principal *entities.Auth) wash_servers.ListResponder {
			return wash_servers.ListNotImplemented()
		})
	}
	if api.WashServersNotificationHandler == nil {
		api.WashServersNotificationHandler = wash_servers.NotificationHandlerFunc(func(params wash_servers.NotificationParams, principal *entities.Auth) wash_servers.NotificationResponder {
			return wash_servers.NotificationNotImplemented()
		})
	}
	if api.WashServersUpdateHandler == nil {
		api.WashServersUpdateHandler = wash_servers.UpdateHandlerFunc(func(params wash_servers.UpdateParams, principal *entities.Auth) wash_servers.UpdateResponder {
			return wash_servers.UpdateNotImplemented()
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
