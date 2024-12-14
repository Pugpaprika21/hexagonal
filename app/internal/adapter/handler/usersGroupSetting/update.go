package usersgroupsetting

import (
	"net/http"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/constant"
	"pugpaprika/app/pkg/respone"

	"github.com/labstack/echo/v4"
)

func (u *usersGroupSettingHandler) UpdateUsersGroupSetting(c echo.Context) error {
	var req request.UpdateUsersGroupSetting
	var resp = respone.NewBuilder()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Code(constant.FOR_BAD_REQUEST).Message(err.Error()).Build())
	}

	err := u.serv.UpdateUsersGroupSetting(c.Request().Context(), req.UpdateUsersGroupSettingRows)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.Code(constant.FOR_ERROR).Message(err.Error()).Build())
	}

	return c.JSON(http.StatusOK, resp.Code(constant.FOR_SUCCESS).Message("success").Build())
}
