package factory

import (
	"errors"
	"pugpaprika/app/pkg/database"

	"gorm.io/gorm"
)

type IDbFactory interface {
	Connect() (*gorm.DB, error)
	Close() error
}

func NewDatabase(driver string) (IDbFactory, error) {
	switch driver {
	case "pgsql":
		return database.NewPgSqlConnector(), nil
	case "mysql":
		return database.NewMySqlConnector(), nil
	default:
		return nil, errors.New("unsupported database type")
	}
}
