package users

import (
	"net/http"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/constant"
	"pugpaprika/app/pkg/respone"

	"github.com/labstack/echo/v4"
)

func (u *userHandler) DeleteUsers(c echo.Context) error {
	var req request.DeleteUsers
	var resp = respone.NewBuilder()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Code(constant.FOR_BAD_REQUEST).Message(err.Error()).Build())
	}

	err := u.serv.DeleteUsers(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.Code(constant.FOR_ERROR).Message(err.Error()).Build())
	}

	return c.JSON(http.StatusOK, resp.Code(constant.FOR_SUCCESS).Message("success").Build())
}
