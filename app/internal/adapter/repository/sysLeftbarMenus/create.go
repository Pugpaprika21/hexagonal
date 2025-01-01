package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/internal/core/schema"
)

func (s *sysLeftBarMenusRepository) CreateMenus(ctx context.Context, parmObj []schema.CreateMenus) error {
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table("sys_leftbar_menus").Create(&parmObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
