package repository

import (
	users "pugpaprika/app/internal/adapter/repository/users"
	usersgroup "pugpaprika/app/internal/adapter/repository/usersGroup"
	usersgroupsetting "pugpaprika/app/internal/adapter/repository/usersGroupSetting"

	"gorm.io/gorm"
)

type Repository struct {
	Users             users.IUsersRepository
	UsersGroup        usersgroup.IUsersGroupRepository
	UsersGroupSetting usersgroupsetting.IUsersGroupSettingRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users:             users.NewUsersRepository(db),
		UsersGroup:        usersgroup.NewUsersGroupRepository(db),
		UsersGroupSetting: usersgroupsetting.NewUsersGroupSettingRepository(db),
	}
}
