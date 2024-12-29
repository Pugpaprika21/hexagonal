package sysleftbarmenus

import (
	sysleftbarmenus "pugpaprika/app/internal/core/service/sysLeftbarMenus"

	"github.com/labstack/echo/v4"
)

type ISyLeftBarMenusHandler interface {
	GetMainManus(c echo.Context) error
}

type syLeftBarMenusHandler struct {
	serv sysleftbarmenus.ISysLeftBarMenusService
}

func NewSyLeftBarMenusHandler(serv sysleftbarmenus.ISysLeftBarMenusService) ISyLeftBarMenusHandler {
	return &syLeftBarMenusHandler{serv: serv}
}
