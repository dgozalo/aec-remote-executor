package database

import (
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// PostgresDBAccess is the database access object
type PostgresDBAccess struct {
	DB *gorm.DB
}

// NewPostgresDBAccess creates a new database access object and initializes the connection to the database
func NewPostgresDBAccess() (*PostgresDBAccess, error) {
	dbURL := os.Getenv("POSTGRES_DB_CONNECT_STRING")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/aec_executor_dev?sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, errors.Wrap(err, "there was a problem connecting to the database")
	}
	return &PostgresDBAccess{
		DB: db,
	}, nil
}
