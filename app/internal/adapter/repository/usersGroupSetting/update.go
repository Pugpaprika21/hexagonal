package usersgroupsetting

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (u *UsersGroupSettingRepository) UpdateUsersGroupSetting(ctx context.Context, parmObj []schema.UpdateUsersGroupSetting, sql sqlx.Sqlx) error {
	tx := u.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table("users_group_setting").Where(sql.Stmt, sql.Args...).Updates(&parmObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
