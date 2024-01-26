package handlers

import (
	"sbp/internal/app"

	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.SugaredLogger
	svc    app.Service
}

func NewHandler(logger *zap.SugaredLogger, svc app.Service) Handler {
	return Handler{
		logger: logger,
		svc:    svc,
	}
}
