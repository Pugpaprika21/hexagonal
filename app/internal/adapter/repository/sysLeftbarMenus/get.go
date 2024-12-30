package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (s *sysLeftBarMenusRepository) GetMainMenus(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetNewMainManus, error) {
	var rows []schema.GetNewMainManus
	result := s.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}

func (s *sysLeftBarMenusRepository) GetParentMenus(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetParentMenus, error) {
	var rows []schema.GetParentMenus
	result := s.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}

func (s *sysLeftBarMenusRepository) GetChildMenus(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetChildMenus, error) {
	var rows []schema.GetChildMenus
	result := s.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}
