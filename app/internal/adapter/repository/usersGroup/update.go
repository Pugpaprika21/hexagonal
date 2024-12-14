package usersgroup

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (u *usersGroupRepository) UpdateUsersGroup(ctx context.Context, parmObj []schema.UpdateUsersGroup, sql sqlx.Sqlx) error {
	tx := u.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table("users_group").Where(sql.Stmt, sql.Args...).Updates(&parmObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
