package users

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (u *userRepository) UpdateUsers(ctx context.Context, parmObj []schema.UpdateUsers, sql sqlx.Sqlx) error {
	tx := u.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table("users").Where(sql.Stmt, sql.Args...).Updates(&parmObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
