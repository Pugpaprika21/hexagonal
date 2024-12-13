package users

import (
	"pugpaprika/app/internal/core/service/users"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	GetUsers(c echo.Context) error
}

type userHandler struct {
	users users.IUsersSevice
}

func NewUserHandler(users users.IUsersSevice) IUserHandler {
	return &userHandler{users: users}
}
