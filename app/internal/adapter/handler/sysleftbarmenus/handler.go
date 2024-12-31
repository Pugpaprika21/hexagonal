package sysleftbarmenus

import (
	sysleftbarmenus "pugpaprika/app/internal/core/service/sysLeftbarMenus"

	"github.com/labstack/echo/v4"
)

type ISyLeftBarMenusHandler interface {
	GetMainMenus(c echo.Context) error
	GetAllMenus(c echo.Context) error
	CreateMenus(c echo.Context) error
	UpdateMenus(c echo.Context) error
	DeleteMenus(c echo.Context) error
}

type syLeftBarMenusHandler struct {
	serv sysleftbarmenus.ISysLeftBarMenusService
}

func NewSyLeftBarMenusHandler(serv sysleftbarmenus.ISysLeftBarMenusService) ISyLeftBarMenusHandler {
	return &syLeftBarMenusHandler{serv: serv}
}
