package repository

import (
	"errors"
	logic "sbp/internal/logic"
	"sbp/pkg/bootstrap"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

// check 'Repository' struct = logic interface 'Repository'
var _ = logic.Repository(&Repository{})

const layer = "repository"

// Repository ...
type Repository struct {
	db *dbr.Connection
	l  *zap.SugaredLogger
}

// RepositoryConfi[g ...
type RepositoryConfig struct {
	DBConfig *bootstrap.DBConfig
	Logger   *zap.SugaredLogger
}

// checkRepositoryConfig ...
func checkRepositoryConfig(conf RepositoryConfig) error {
	if conf.DBConfig == nil {
		return errors.New("repository db config is empty")
	}
	if conf.Logger == nil {
		return errors.New("repository logger is empty")
	}

	return nil
}

// NewRepository ...
func NewRepository(repositoryConfig RepositoryConfig) (*Repository, error) {
	dbConfig := *repositoryConfig.DBConfig
	logger := repositoryConfig.Logger

	// check repository config
	err := checkRepositoryConfig(repositoryConfig)
	if err != nil {
		return nil, bootstrap.CustomError(layer, "checkRepository", err)
	}

	// new db conn
	dbConn, err := bootstrap.NewDbConn(dbConfig)
	if err != nil {
		return nil, bootstrap.CustomError(layer, "new db conn: %s", err)
	}
	repositoryConfig.Logger.Debug("connected to db")

	// up migrations
	err = bootstrap.UpMigrations(dbConn.DB, dbConfig.Database, "internal/repository/migrations")
	if err != nil {
		return nil, bootstrap.CustomError(layer, "up migrations: %s", err)
	}
	logger.Debug("applied migrations")

	return &Repository{
		db: dbConn,
		l:  logger,
	}, nil
}

// Close ...
func (rep *Repository) Close() error {
	return rep.db.Close()
}
