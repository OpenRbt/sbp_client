package bootstrap

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	levelDev     = "dev"
	levelInfo    = "info"
	levelWarning = "warning"
	levelError   = "error"
	levelDefault = levelDev
)

// knownLevels ...
var knownLevels map[string]bool = map[string]bool{
	levelDev:     true,
	levelInfo:    true,
	levelWarning: true,
	levelError:   true,
}

func NewLogger(level string) (l *zap.SugaredLogger, err error) {

	sLevel := strings.ToLower(level)
	_, ok := knownLevels[sLevel]
	if !ok {
		fmt.Printf("unkown logger level '%s', default logger '%s' is used\n", level, levelDefault)
		sLevel = levelDefault
	}
	cfg := zap.NewProductionConfig()
	cfg.DisableCaller = true

	switch sLevel {
	default:
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	case levelDev:
		fmt.Printf("logger '%s' is used\n", levelDev)
		cfg = zap.NewDevelopmentConfig()
		cfg.DisableCaller = true
		cfg.DisableStacktrace = true
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	case levelInfo:
		fmt.Printf("logger '%s' is used\n", levelInfo)
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case levelWarning:
		fmt.Printf("logger '%s' is used\n", levelWarning)
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case levelError:
		fmt.Printf("logger '%s' is used\n", levelError)
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	basicLogger, err := cfg.Build()
	l = basicLogger.Sugar()

	return
}
