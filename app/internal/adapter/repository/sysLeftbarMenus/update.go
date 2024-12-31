package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (s *sysLeftBarMenusRepository) UpdateMenus(ctx context.Context, parmObj []schema.UpdateMenus, sql sqlx.Sqlx) error {
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table("sys_leftbar_menus").Where(sql.Stmt, sql.Args...).Updates(&parmObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
