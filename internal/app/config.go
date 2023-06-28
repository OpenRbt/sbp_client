package app

import (
	"fmt"
	"sbp/pkg/bootstrap"

	"github.com/joho/godotenv"
)

// getConfig ...
func getConfig(envFilePath string) (*bootstrap.Config, error) {
	// read envs
	err := godotenv.Load(envFilePath)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %s", err.Error())
	}

	config, err := bootstrap.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("new config: %s", err.Error())
	}

	return config, nil
}
