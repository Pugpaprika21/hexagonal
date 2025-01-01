package mstparamshdr

import (
	"net/http"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/constant"
	"pugpaprika/app/pkg/response"

	"github.com/labstack/echo/v4"
)

func (m *mstParamsHdrHandler) DeleteParamsHdr(c echo.Context) error {
	var req request.DeleteParamsHdr
	var resp = response.NewBuilder()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Code(constant.FOR_BAD_REQUEST).Message(err.Error()).Build())
	}

	err := m.serv.DeleteParamsHdr(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.Code(constant.FOR_ERROR).Message(err.Error()).Build())
	}

	return c.JSON(http.StatusOK, resp.Code(constant.FOR_SUCCESS).Message("success").Build())
}
