package repository

import (
	"pugpaprika/app/internal/adapter/repository/migrations"
	users "pugpaprika/app/internal/adapter/repository/users"
	usersgroup "pugpaprika/app/internal/adapter/repository/usersGroup"
	usersgroupsetting "pugpaprika/app/internal/adapter/repository/usersGroupSetting"

	"gorm.io/gorm"
)

type Repository struct {
	Users             users.IUsersRepository
	UsersGroup        usersgroup.IUsersGroupRepository
	UsersGroupSetting usersgroupsetting.IUsersGroupSettingRepository
	Migrations        migrations.IMigrationsRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users:             users.NewUsersRepository(db),
		UsersGroup:        usersgroup.NewUsersGroupRepository(db),
		UsersGroupSetting: usersgroupsetting.NewUsersGroupSettingRepository(db),
		Migrations:        migrations.NewMigrationsRepository(db),
	}
}
