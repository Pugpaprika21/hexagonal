package usersgroupsetting

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"

	"gorm.io/gorm"
)

type IUsersGroupSettingRepository interface {
	GetUsersGroupSetting(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetUsersGroupSetting, error)
	GetUsersGroupSettingByUserID(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetUsersGroupSettingByUserID, error)
	CreateUsersGroupSetting(ctx context.Context, parmObj []schema.CreateUsersGroupSetting) error
	UpdateUsersGroupSetting(ctx context.Context, parmObj []schema.UpdateUsersGroupSetting, sql sqlx.Sqlx) error
	DeleteUsersSettingGroup(ctx context.Context, sql sqlx.Sqlx) error
}

type UsersGroupSettingRepository struct {
	db *gorm.DB
}

func NewUsersGroupSettingRepository(db *gorm.DB) IUsersGroupSettingRepository {
	return &UsersGroupSettingRepository{db: db}
}
