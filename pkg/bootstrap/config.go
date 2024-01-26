package bootstrap

import (
	"errors"
	"os"
	"sbp/internal/config"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func NewConfig(configFiles ...string) (*config.Config, error) {
	var c config.Config
	err := godotenv.Load(configFiles...)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	return &c, env.Parse(&c, env.Options{
		RequiredIfNoDef: true,
	})

}
