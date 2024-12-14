package usersgroup

import (
	usersgroup "pugpaprika/app/internal/core/service/usersGroup"

	"github.com/labstack/echo/v4"
)

type IUsersGroupHandler interface {
	GetUsersGroup(c echo.Context) error
	CreateUsersGroup(c echo.Context) error
	UpdateUsersGroup(c echo.Context) error
	DeleteUsersGroup(c echo.Context) error
}

type usersGroupHandler struct {
	serv usersgroup.IUsersGroupService
}

func NewUsersGroupHandler(serv usersgroup.IUsersGroupService) IUsersGroupHandler {
	return &usersGroupHandler{serv: serv}
}
