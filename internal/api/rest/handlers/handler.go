package handlers

import "go.uber.org/zap"

// Handler ...
type Handler struct {
	logger *zap.SugaredLogger
	logic  Logic
}

// NewHandler ...
func NewHandler(logger *zap.SugaredLogger, logic Logic) Handler {
	return Handler{
		logger: logger,
		logic:  logic,
	}
}
