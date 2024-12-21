package migrations

import (
	"net/http"
	"pugpaprika/app/pkg/constant"
	"pugpaprika/app/pkg/response"

	"github.com/labstack/echo/v4"
)

func (m *migrationsHandler) CreateGoStructProcedure(c echo.Context) error {
	var resp = response.NewBuilder()

	err := m.serv.CreateGoStructProcedure(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.Code(constant.FOR_ERROR).Message(err.Error()).Build())
	}

	return c.JSON(http.StatusOK, resp.Code(constant.FOR_SUCCESS).Message("success").Build())
}
