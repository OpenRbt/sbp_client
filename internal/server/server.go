package server

import (
	"net/http"
	swaggerRestapi "sbp/internal/openapi/restapi"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

// Server ...
type Server struct {
	swaggerServer *swaggerRestapi.Server
	logger        *zap.SugaredLogger
}

// ServerConfig ...
type ServerConfig struct {
	Logger         *zap.SugaredLogger
	Host           string
	Port           string
	AllowedOrigins string
	Api            Api
}

// checkServerConfig ...
func checkServerConfig(conf ServerConfig) error {
	if conf.Host == "" {
		return errors.New("server host is empty")
	}
	if conf.Port == "" {
		return errors.New("server port is empty")
	}
	if conf.AllowedOrigins == "" {
		return errors.New("server allowedOrigins is empty")
	}
	if conf.Api == nil {
		return errors.New("server api is empty")
	}
	return nil
}

// NewServer ...
func NewServer(config ServerConfig) (*Server, error) {
	// check server config
	err := checkServerConfig(config)
	if err != nil {
		return nil, err
	}

	// init swagger server
	server := swaggerRestapi.NewServer(config.Api.GetSwaggerApi())
	server.Host = string(config.Host)
	portInt, err := strconv.ParseInt(config.Port, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse port")
	}
	server.Port = int(portInt)

	// set cors
	newCORS := cors.New(cors.Options{
		// AllowedOrigins:   splitCommaSeparatedStr(config.AllowedOrigins),
		AllowedMethods:   []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})
	// newCORS.Log = cors.Logger(structlog.New(structlog.KeyUnit, "CORS"))

	// fix empty origin
	emptyOriginHandler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Origin") == "" {
				r.Header.Add("Origin", "empty")
			}
			h.ServeHTTP(w, r)
		})
	}
	server.SetHandler(emptyOriginHandler(newCORS.Handler(config.Api.GetHandler())))

	return &Server{
		swaggerServer: server,
		logger:        config.Logger,
	}, nil
}

// Run ...
func (s *Server) Run() error {
	s.logger.Info("started server at:", s.swaggerServer.Port)
	return s.swaggerServer.Serve()

}

// splitCommaSeparatedStr ...
func splitCommaSeparatedStr(commaSeparated string) (result []string) {
	for _, item := range strings.Split(commaSeparated, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			result = append(result, item)
		}
	}
	return result
}
