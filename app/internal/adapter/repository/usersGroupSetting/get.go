package usersgroupsetting

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (u *usersGroupSettingRepository) GetUsersGroupSetting(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetUsersGroupSetting, error) {
	var rows []schema.GetUsersGroupSetting
	result := u.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}

func (u *usersGroupSettingRepository) GetUsersGroupSettingByUserID(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetUsersGroupSettingByUserID, error) {
	var rows []schema.GetUsersGroupSettingByUserID
	result := u.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}
