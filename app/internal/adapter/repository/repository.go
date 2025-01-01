package repository

import (
	"pugpaprika/app/internal/adapter/repository/migrations"
	mstparamsdtl "pugpaprika/app/internal/adapter/repository/mstParamsDtl"
	mstparamshdr "pugpaprika/app/internal/adapter/repository/mstParamsHdr"
	sysleftbarmenus "pugpaprika/app/internal/adapter/repository/sysLeftbarMenus"
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
	SysLeftBarMenus   sysleftbarmenus.ISysLeftBarMenusRepository
	MstParamsHdr      mstparamshdr.IMstParamsHdrRepository
	MstParamsDtl      mstparamsdtl.IMstParamsDtlRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users:             users.NewUsersRepository(db),
		UsersGroup:        usersgroup.NewUsersGroupRepository(db),
		UsersGroupSetting: usersgroupsetting.NewUsersGroupSettingRepository(db),
		Migrations:        migrations.NewMigrationsRepository(db),
		SysLeftBarMenus:   sysleftbarmenus.NewSysLeftBarMenusRepository(db),
		MstParamsHdr:      mstparamshdr.NewMstParamsHdrRepository(db),
		MstParamsDtl:      mstparamsdtl.NewMstParamsDtlRepository(db),
	}
}
