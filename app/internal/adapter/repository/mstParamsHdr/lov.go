package mstparamshdr

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (m *mstParamsHdrRepository) LovParamsHdr(ctx context.Context, sql sqlx.Sqlx) ([]schema.LovParamsHdr, error) {
	var rows []schema.LovParamsHdr
	result := m.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}
