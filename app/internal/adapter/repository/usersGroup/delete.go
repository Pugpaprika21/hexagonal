package usersgroup

import (
	"context"
	"pugpaprika/app/pkg/sqlx"
)

func (u *usersGroupRepository) DeleteUsersGroup(ctx context.Context, sql sqlx.Sqlx) error {
	tx := u.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	result := tx.Raw(sql.Stmt, sql.Args...)
	if result.Error != nil {
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
