package rest

import (
	"net/http"
	"sbp/internal/config"
	swaggerRestapi "sbp/openapi/restapi"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

type Server struct {
	swaggerServer *swaggerRestapi.Server
	logger        *zap.SugaredLogger
}

func NewServer(cfg config.ServerConfig) (*Server, error) {
	err := cfg.CheckServerConfig()
	if err != nil {
		return nil, err
	}

	server := swaggerRestapi.NewServer(cfg.Api.GetSwaggerApi())
	server.Host = string(cfg.Host)
	portInt, err := strconv.ParseInt(cfg.Port, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse port")
	}
	server.Port = int(portInt)

	newCORS := cors.New(cors.Options{
		AllowedMethods:   []string{"POST", "PATCH", "PUT", "GET", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	emptyOriginHandler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Origin") == "" {
				r.Header.Add("Origin", "empty")
			}
			h.ServeHTTP(w, r)
		})
	}
	server.SetHandler(emptyOriginHandler(newCORS.Handler(cfg.Api.GetHandler())))

	return &Server{
		swaggerServer: server,
		logger:        cfg.Logger,
	}, nil
}

func (s *Server) Run() error {
	s.logger.Info("started server at:", s.swaggerServer.Port)
	return s.swaggerServer.Serve()

}
