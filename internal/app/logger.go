package app

import (
	"fmt"
	"sbp/pkg/bootstrap"

	"go.uber.org/zap"
)

// getLogger ...
func getLogger(logLevel string) (l *zap.SugaredLogger, err error) {
	logger, err := bootstrap.NewLogger(logLevel)
	if err != nil {
		return nil, fmt.Errorf("new logger: %s", err.Error())
	}
	return logger, nil
}
