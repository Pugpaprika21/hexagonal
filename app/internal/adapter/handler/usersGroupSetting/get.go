package usersgroupsetting

import (
	"net/http"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/constant"
	"pugpaprika/app/pkg/response"

	"github.com/labstack/echo/v4"
)

func (u *usersGroupSettingHandler) GetUsersGroupSetting(c echo.Context) error {
	var req request.GetUsersGroupSetting
	var resp = response.NewBuilder()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Code(constant.FOR_BAD_REQUEST).Message(err.Error()).Build())
	}

	data, err := u.serv.GetUsersGroupSetting(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.Code(constant.FOR_ERROR).Message(err.Error()).Build())
	}

	return c.JSON(http.StatusOK, resp.Code(constant.FOR_SUCCESS).Message("success").Data(data).Build())
}

func (u *usersGroupSettingHandler) GetUsersGroupSettingByUserID(c echo.Context) error {
	var req request.GetUsersGroupSettingByUserID
	var resp = response.NewBuilder()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Code(constant.FOR_BAD_REQUEST).Message(err.Error()).Build())
	}

	data, err := u.serv.GetUsersGroupSettingByUserID(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.Code(constant.FOR_ERROR).Message(err.Error()).Build())
	}

	return c.JSON(http.StatusOK, resp.Code(constant.FOR_SUCCESS).Message("success").Data(data).Build())
}
