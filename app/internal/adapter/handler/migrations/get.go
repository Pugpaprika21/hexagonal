package migrations

import (
	"net/http"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/constant"
	"pugpaprika/app/pkg/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func (m *migrationsHandler) GetGoStructProcedure(c echo.Context) error {
	var req request.GetGoStructProcedure
	var resp = response.NewBuilder()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Code(constant.FOR_BAD_REQUEST).Message(err.Error()).Build())
	}

	data, err := m.serv.GetGoStructProcedure(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.Code(constant.FOR_ERROR).Message(err.Error()).Build())
	}

	return c.String(http.StatusOK, data)
}

func (m *migrationsHandler) GetAllGoStructProcedure(c echo.Context) error {
	var req request.GetAllGoStructProcedure
	var resp = response.NewBuilder()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Code(constant.FOR_BAD_REQUEST).Message(err.Error()).Build())
	}

	data, err := m.serv.GetAllGoStructProcedure(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.Code(constant.FOR_ERROR).Message(err.Error()).Build())
	}

	return c.String(http.StatusOK, strings.Join(data, " "))
}
