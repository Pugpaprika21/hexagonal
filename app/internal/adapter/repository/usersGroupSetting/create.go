package usersgroupsetting

import (
	"context"
	"pugpaprika/app/internal/core/schema"
)

func (u *UsersGroupSettingRepository) CreateUsersGroupSetting(ctx context.Context, parmObj []schema.CreateUsersGroupSetting) error {
	tx := u.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table("users_group_setting").Create(&parmObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
