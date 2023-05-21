package database

import (
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type PostgresDBAccess struct {
	DB *gorm.DB
}

func NewPostgresDBAccess() (*PostgresDBAccess, error) {
	dbURL := os.Getenv("POSTGRES_DB_CONNECT_STRING")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/aec_executor_dev?sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "there was a problem connecting to the database")
	}
	return &PostgresDBAccess{
		DB: db,
	}, nil
}
