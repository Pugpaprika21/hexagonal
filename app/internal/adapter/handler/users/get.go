package users

import (
	"net/http"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/respone"

	"github.com/labstack/echo/v4"
)

func (u *userHandler) GetUsers(c echo.Context) error {
	var req request.GetUsers
	var resp = respone.NewBuilder()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Code(400).Message(err.Error()).Build())
	}

	data, err := u.users.GetUsers(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.Code(500).Message(err.Error()).Build())
	}

	return c.JSON(http.StatusOK, resp.Code(1000).Message("success").Data(data).Build())
}
