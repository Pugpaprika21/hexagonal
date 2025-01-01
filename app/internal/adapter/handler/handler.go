package handler

import (
	migrations "pugpaprika/app/internal/adapter/handler/migrations"
	mstparamsdtl "pugpaprika/app/internal/adapter/handler/mstParamsDtl"
	mstparamshdr "pugpaprika/app/internal/adapter/handler/mstParamsHdr"
	sysleftbarmenus "pugpaprika/app/internal/adapter/handler/sysleftbarmenus"
	users "pugpaprika/app/internal/adapter/handler/users"
	usersgroup "pugpaprika/app/internal/adapter/handler/usersGroup"
	usersgroupsetting "pugpaprika/app/internal/adapter/handler/usersGroupSetting"
	"pugpaprika/app/internal/core/service"
)

type Handler struct {
	Users             users.IUserHandler
	UsersGroup        usersgroup.IUsersGroupHandler
	UsersGroupSetting usersgroupsetting.IUsersGroupSettingHandler
	Migrations        migrations.IMigrationsHandler
	SysLeftBarMenus   sysleftbarmenus.ISyLeftBarMenusHandler
	MstParamsHdr      mstparamshdr.IMstParamsHdrHandler
	MstParamsDtl      mstparamsdtl.IMstParamsDtlHandler
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		Users:             users.NewUserHandler(serv.Users),
		UsersGroup:        usersgroup.NewUsersGroupHandler(serv.UsersGroup),
		UsersGroupSetting: usersgroupsetting.NewUsersGroupSettingHandler(serv.UsersGroupSetting),
		Migrations:        migrations.NewMigrationsHandler(serv.Migrations),
		SysLeftBarMenus:   sysleftbarmenus.NewSyLeftBarMenusHandler(serv.SysLeftBarMenus),
		MstParamsHdr:      mstparamshdr.NewMstParamsHdrHandler(serv.MstParamsHdr),
		MstParamsDtl:      mstparamsdtl.NewMstParamsDtlHandler(serv.MstParamsDtl),
	}
}
