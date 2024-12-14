package users

import (
	"pugpaprika/app/internal/core/service/users"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	GetUsers(c echo.Context) error
	CreateUsers(c echo.Context) error
	UpdateUsers(c echo.Context) error
	DeleteUsers(c echo.Context) error
}

type userHandler struct {
	serv users.IUsersSevice
}

func NewUserHandler(serv users.IUsersSevice) IUserHandler {
	return &userHandler{serv: serv}
}
