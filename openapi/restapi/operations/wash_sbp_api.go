// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"sbp/internal/entities"
	"sbp/openapi/restapi/operations/notifications"
	"sbp/openapi/restapi/operations/payments"
	"sbp/openapi/restapi/operations/standard"
	"sbp/openapi/restapi/operations/transactions"
	"sbp/openapi/restapi/operations/washes"
)

// NewWashSbpAPI creates a new WashSbp instance
func NewWashSbpAPI(spec *loads.Document) *WashSbpAPI {
	return &WashSbpAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		WashesAssignWashToGroupHandler: washes.AssignWashToGroupHandlerFunc(func(params washes.AssignWashToGroupParams, principal *entities.Auth) washes.AssignWashToGroupResponder {
			return washes.AssignWashToGroupNotImplemented()
		}),
		PaymentsCancelPaymentHandler: payments.CancelPaymentHandlerFunc(func(params payments.CancelPaymentParams, principal *entities.Auth) payments.CancelPaymentResponder {
			return payments.CancelPaymentNotImplemented()
		}),
		WashesCreateWashHandler: washes.CreateWashHandlerFunc(func(params washes.CreateWashParams, principal *entities.Auth) washes.CreateWashResponder {
			return washes.CreateWashNotImplemented()
		}),
		WashesDeleteWashHandler: washes.DeleteWashHandlerFunc(func(params washes.DeleteWashParams, principal *entities.Auth) washes.DeleteWashResponder {
			return washes.DeleteWashNotImplemented()
		}),
		TransactionsGetTransactionsHandler: transactions.GetTransactionsHandlerFunc(func(params transactions.GetTransactionsParams, principal *entities.Auth) transactions.GetTransactionsResponder {
			return transactions.GetTransactionsNotImplemented()
		}),
		WashesGetWashByIDHandler: washes.GetWashByIDHandlerFunc(func(params washes.GetWashByIDParams, principal *entities.Auth) washes.GetWashByIDResponder {
			return washes.GetWashByIDNotImplemented()
		}),
		WashesGetWashesHandler: washes.GetWashesHandlerFunc(func(params washes.GetWashesParams, principal *entities.Auth) washes.GetWashesResponder {
			return washes.GetWashesNotImplemented()
		}),
		StandardHealthcheckHandler: standard.HealthcheckHandlerFunc(func(params standard.HealthcheckParams, principal *entities.Auth) standard.HealthcheckResponder {
			return standard.HealthcheckNotImplemented()
		}),
		PaymentsInitPaymentHandler: payments.InitPaymentHandlerFunc(func(params payments.InitPaymentParams, principal *entities.Auth) payments.InitPaymentResponder {
			return payments.InitPaymentNotImplemented()
		}),
		NotificationsReceiveNotificationHandler: notifications.ReceiveNotificationHandlerFunc(func(params notifications.ReceiveNotificationParams, principal *entities.Auth) notifications.ReceiveNotificationResponder {
			return notifications.ReceiveNotificationNotImplemented()
		}),
		WashesUpdateWashHandler: washes.UpdateWashHandlerFunc(func(params washes.UpdateWashParams, principal *entities.Auth) washes.UpdateWashResponder {
			return washes.UpdateWashNotImplemented()
		}),

		// Applies when the "Authorization" header is set
		AuthKeyAuth: func(token string) (*entities.Auth, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*WashSbpAPI microservice for the sbp system of self-service car washes */
type WashSbpAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AuthKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	AuthKeyAuth func(string) (*entities.Auth, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// WashesAssignWashToGroupHandler sets the operation handler for the assign wash to group operation
	WashesAssignWashToGroupHandler washes.AssignWashToGroupHandler
	// PaymentsCancelPaymentHandler sets the operation handler for the cancel payment operation
	PaymentsCancelPaymentHandler payments.CancelPaymentHandler
	// WashesCreateWashHandler sets the operation handler for the create wash operation
	WashesCreateWashHandler washes.CreateWashHandler
	// WashesDeleteWashHandler sets the operation handler for the delete wash operation
	WashesDeleteWashHandler washes.DeleteWashHandler
	// TransactionsGetTransactionsHandler sets the operation handler for the get transactions operation
	TransactionsGetTransactionsHandler transactions.GetTransactionsHandler
	// WashesGetWashByIDHandler sets the operation handler for the get wash by Id operation
	WashesGetWashByIDHandler washes.GetWashByIDHandler
	// WashesGetWashesHandler sets the operation handler for the get washes operation
	WashesGetWashesHandler washes.GetWashesHandler
	// StandardHealthcheckHandler sets the operation handler for the healthcheck operation
	StandardHealthcheckHandler standard.HealthcheckHandler
	// PaymentsInitPaymentHandler sets the operation handler for the init payment operation
	PaymentsInitPaymentHandler payments.InitPaymentHandler
	// NotificationsReceiveNotificationHandler sets the operation handler for the receive notification operation
	NotificationsReceiveNotificationHandler notifications.ReceiveNotificationHandler
	// WashesUpdateWashHandler sets the operation handler for the update wash operation
	WashesUpdateWashHandler washes.UpdateWashHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *WashSbpAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *WashSbpAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *WashSbpAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WashSbpAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *WashSbpAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *WashSbpAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WashSbpAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WashSbpAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WashSbpAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WashSbpAPI
func (o *WashSbpAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AuthKeyAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.WashesAssignWashToGroupHandler == nil {
		unregistered = append(unregistered, "washes.AssignWashToGroupHandler")
	}
	if o.PaymentsCancelPaymentHandler == nil {
		unregistered = append(unregistered, "payments.CancelPaymentHandler")
	}
	if o.WashesCreateWashHandler == nil {
		unregistered = append(unregistered, "washes.CreateWashHandler")
	}
	if o.WashesDeleteWashHandler == nil {
		unregistered = append(unregistered, "washes.DeleteWashHandler")
	}
	if o.TransactionsGetTransactionsHandler == nil {
		unregistered = append(unregistered, "transactions.GetTransactionsHandler")
	}
	if o.WashesGetWashByIDHandler == nil {
		unregistered = append(unregistered, "washes.GetWashByIDHandler")
	}
	if o.WashesGetWashesHandler == nil {
		unregistered = append(unregistered, "washes.GetWashesHandler")
	}
	if o.StandardHealthcheckHandler == nil {
		unregistered = append(unregistered, "standard.HealthcheckHandler")
	}
	if o.PaymentsInitPaymentHandler == nil {
		unregistered = append(unregistered, "payments.InitPaymentHandler")
	}
	if o.NotificationsReceiveNotificationHandler == nil {
		unregistered = append(unregistered, "notifications.ReceiveNotificationHandler")
	}
	if o.WashesUpdateWashHandler == nil {
		unregistered = append(unregistered, "washes.UpdateWashHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WashSbpAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WashSbpAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "authKey":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.AuthKeyAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *WashSbpAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *WashSbpAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *WashSbpAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *WashSbpAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the wash sbp API
func (o *WashSbpAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *WashSbpAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/groups/{groupId}/washes/{washId}"] = washes.NewAssignWashToGroup(o.context, o.WashesAssignWashToGroupHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/payments/cancel"] = payments.NewCancelPayment(o.context, o.PaymentsCancelPaymentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/washes"] = washes.NewCreateWash(o.context, o.WashesCreateWashHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/washes/{id}"] = washes.NewDeleteWash(o.context, o.WashesDeleteWashHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/transactions"] = transactions.NewGetTransactions(o.context, o.TransactionsGetTransactionsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/washes/{id}"] = washes.NewGetWashByID(o.context, o.WashesGetWashByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/washes"] = washes.NewGetWashes(o.context, o.WashesGetWashesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthcheck"] = standard.NewHealthcheck(o.context, o.StandardHealthcheckHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/payments/init"] = payments.NewInitPayment(o.context, o.PaymentsInitPaymentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/notification"] = notifications.NewReceiveNotification(o.context, o.NotificationsReceiveNotificationHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/washes/{id}"] = washes.NewUpdateWash(o.context, o.WashesUpdateWashHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WashSbpAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *WashSbpAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *WashSbpAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *WashSbpAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *WashSbpAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
