package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (s *sysLeftBarMenusRepository) GetMainManus(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetNewMainManus, error) {
	var rows []schema.GetNewMainManus
	result := s.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}
