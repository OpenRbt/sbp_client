package app

import (
	"golang.org/x/net/context"
)

// App ...
type App struct {
	deps *deps
}

// NewApp ...
func NewApp(ctx context.Context, envFilePath string) (*App, error) {
	config, err := getConfig(envFilePath)
	if err != nil {
		return nil, err
	}

	logger, err := getLogger(config.LogLevel)
	if err != nil {
		return nil, err
	}

	deps, err := InitDeps(ctx, config, logger)
	if err != nil {
		return nil, err
	}

	return &App{
		deps: deps,
	}, nil
}

// Run ...
func (a App) Run() error {
	return a.deps.httpServer.Run()
}

// Run ...
func (a App) Close() {
	a.deps.close()
}
