package handler

import (
	users "pugpaprika/app/internal/adapter/handler/users"
	usersgroup "pugpaprika/app/internal/adapter/handler/usersGroup"
	usersgroupsetting "pugpaprika/app/internal/adapter/handler/usersGroupSetting"
	"pugpaprika/app/internal/core/service"
)

type Handler struct {
	Users             users.IUserHandler
	UsersGroup        usersgroup.IUsersGroupHandler
	UsersGroupSetting usersgroupsetting.IUsersGroupSettingHandler
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		Users:             users.NewUserHandler(serv.Users),
		UsersGroup:        usersgroup.NewUsersGroupHandler(serv.UsersGroup),
		UsersGroupSetting: usersgroupsetting.NewUsersGroupSettingHandler(serv.UsersGroupSetting),
	}
}
