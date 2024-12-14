package users

import (
	"context"
	"pugpaprika/app/internal/core/schema"
)

func (u *userRepository) CreateUsers(ctx context.Context, parmObj []schema.CreateUsers) error {
	tx := u.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table("users").Create(&parmObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
