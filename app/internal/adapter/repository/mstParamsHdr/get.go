package mstparamshdr

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (m *mstParamsHdrRepository) GetParamsHdr(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetParamsHdr, error) {
	var rows []schema.GetParamsHdr
	result := m.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}
