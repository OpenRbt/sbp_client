// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"sbp/internal/entities"
	"sbp/openapi/restapi/operations"
	"sbp/openapi/restapi/operations/notifications"
	"sbp/openapi/restapi/operations/payments"
	"sbp/openapi/restapi/operations/standard"
	"sbp/openapi/restapi/operations/washes"
)

//go:generate swagger generate server --target ../../openapi --name WashSbp --spec ../swagger.yaml --principal sbp/internal/entities.Auth --exclude-main --strict-responders

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

	if api.WashesAssignWashToGroupHandler == nil {
		api.WashesAssignWashToGroupHandler = washes.AssignWashToGroupHandlerFunc(func(params washes.AssignWashToGroupParams, principal *entities.Auth) washes.AssignWashToGroupResponder {
			return washes.AssignWashToGroupNotImplemented()
		})
	}
	if api.PaymentsCancelPaymentHandler == nil {
		api.PaymentsCancelPaymentHandler = payments.CancelPaymentHandlerFunc(func(params payments.CancelPaymentParams, principal *entities.Auth) payments.CancelPaymentResponder {
			return payments.CancelPaymentNotImplemented()
		})
	}
	if api.WashesCreateWashHandler == nil {
		api.WashesCreateWashHandler = washes.CreateWashHandlerFunc(func(params washes.CreateWashParams, principal *entities.Auth) washes.CreateWashResponder {
			return washes.CreateWashNotImplemented()
		})
	}
	if api.WashesDeleteWashHandler == nil {
		api.WashesDeleteWashHandler = washes.DeleteWashHandlerFunc(func(params washes.DeleteWashParams, principal *entities.Auth) washes.DeleteWashResponder {
			return washes.DeleteWashNotImplemented()
		})
	}
	if api.WashesGetWashByIDHandler == nil {
		api.WashesGetWashByIDHandler = washes.GetWashByIDHandlerFunc(func(params washes.GetWashByIDParams, principal *entities.Auth) washes.GetWashByIDResponder {
			return washes.GetWashByIDNotImplemented()
		})
	}
	if api.WashesGetWashesHandler == nil {
		api.WashesGetWashesHandler = washes.GetWashesHandlerFunc(func(params washes.GetWashesParams, principal *entities.Auth) washes.GetWashesResponder {
			return washes.GetWashesNotImplemented()
		})
	}
	if api.StandardHealthcheckHandler == nil {
		api.StandardHealthcheckHandler = standard.HealthcheckHandlerFunc(func(params standard.HealthcheckParams, principal *entities.Auth) standard.HealthcheckResponder {
			return standard.HealthcheckNotImplemented()
		})
	}
	if api.PaymentsInitPaymentHandler == nil {
		api.PaymentsInitPaymentHandler = payments.InitPaymentHandlerFunc(func(params payments.InitPaymentParams, principal *entities.Auth) payments.InitPaymentResponder {
			return payments.InitPaymentNotImplemented()
		})
	}
	if api.NotificationsReceiveNotificationHandler == nil {
		api.NotificationsReceiveNotificationHandler = notifications.ReceiveNotificationHandlerFunc(func(params notifications.ReceiveNotificationParams, principal *entities.Auth) notifications.ReceiveNotificationResponder {
			return notifications.ReceiveNotificationNotImplemented()
		})
	}
	if api.WashesUpdateWashHandler == nil {
		api.WashesUpdateWashHandler = washes.UpdateWashHandlerFunc(func(params washes.UpdateWashParams, principal *entities.Auth) washes.UpdateWashResponder {
			return washes.UpdateWashNotImplemented()
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
