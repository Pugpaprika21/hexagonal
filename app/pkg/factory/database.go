package factory

import (
	"errors"
	"os"
	"pugpaprika/app/pkg/database"

	"gorm.io/gorm"
)

type IDbFactory interface {
	Connect() (*gorm.DB, error)
	Close() error
}

func NewDatabase() (IDbFactory, error) {
	driver := os.Getenv("DB_DRIVER")
	switch driver {
	case "pgsql":
		return database.NewPgSqlConnector(), nil
	case "mysql":
		return database.NewMySqlConnector(), nil
	default:
		return nil, errors.New("unsupported database type")
	}
}
