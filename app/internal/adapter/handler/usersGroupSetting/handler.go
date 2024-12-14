package usersgroupsetting

import (
	usersgroupsetting "pugpaprika/app/internal/core/service/usersGroupSetting"

	"github.com/labstack/echo/v4"
)

type IUsersGroupSettingHandler interface {
	GetUsersGroupSetting(c echo.Context) error
	GetUsersGroupSettingByUserID(c echo.Context) error
	CreateUsersGroupSetting(c echo.Context) error
	UpdateUsersGroupSetting(c echo.Context) error
	DeleteUsersSettingGroup(c echo.Context) error
}

type usersGroupSettingHandler struct {
	serv usersgroupsetting.IUsersGroupSettingService
}

func NewUsersGroupSettingHandler(serv usersgroupsetting.IUsersGroupSettingService) IUsersGroupSettingHandler {
	return &usersGroupSettingHandler{serv: serv}
}
