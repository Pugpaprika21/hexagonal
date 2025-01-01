package mstparamsdtl

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (m *mstParamsDtlRepository) LovParamsDtl(ctx context.Context, sql sqlx.Sqlx) ([]schema.LovParamsDtl, error) {
	var rows []schema.LovParamsDtl
	result := m.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}
