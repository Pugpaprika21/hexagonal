package service

import (
	"pugpaprika/app/internal/adapter/repository"
	users "pugpaprika/app/internal/core/service/users"
	usersgroup "pugpaprika/app/internal/core/service/usersGroup"
	usersgroupsetting "pugpaprika/app/internal/core/service/usersGroupSetting"
)

type Service struct {
	Users             users.IUsersSevice
	UsersGroup        usersgroup.IUsersGroupService
	UsersGroupSetting usersgroupsetting.IUsersGroupSettingService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Users:             users.NewUsersService(repos.Users),
		UsersGroup:        usersgroup.NewUsersGroupService(repos.UsersGroup),
		UsersGroupSetting: usersgroupsetting.NewUsersGroupSettingService(repos.UsersGroupSetting),
	}
}
