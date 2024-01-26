package repository

import (
	"sbp/internal/app"
	"sbp/internal/config"
	"sbp/internal/helpers"
	"sbp/pkg/bootstrap"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

var _ = app.Repository(&Repository{})

const layer = "repository"

type Repository struct {
	db *dbr.Connection
	l  *zap.SugaredLogger
}

func NewRepository(cfg config.RepositoryConfig) (*Repository, error) {
	dbConfig := *cfg.DBConfig
	logger := cfg.Logger

	err := cfg.CheckRepositoryConfig()
	if err != nil {
		return nil, helpers.CustomError(layer, "checkRepository", err)
	}

	dbConn, err := bootstrap.NewDbConn(dbConfig)
	if err != nil {
		return nil, helpers.CustomError(layer, "new db conn: %s", err)
	}
	cfg.Logger.Debug("connected to db")

	err = bootstrap.UpMigrations(dbConn.DB, dbConfig.Database, dbConfig.MigrationsPath)
	if err != nil {
		return nil, helpers.CustomError(layer, "up migrations: %s", err)
	}

	logger.Debug("applied migrations")

	return &Repository{
		db: dbConn,
		l:  logger,
	}, nil
}

func (rep *Repository) Close() error {
	return rep.db.Close()
}
