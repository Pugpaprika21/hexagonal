package migrations

import (
	"pugpaprika/app/internal/core/service/migrations"

	"github.com/labstack/echo/v4"
)

type IMigrationsHandler interface {
	GetGoStructProcedure(c echo.Context) error
	CreateGoStructProcedure(c echo.Context) error
}

type migrationsHandler struct {
	serv migrations.IMigrationsService
}

func NewMigrationsHandler(serv migrations.IMigrationsService) IMigrationsHandler {
	return &migrationsHandler{serv: serv}
}
