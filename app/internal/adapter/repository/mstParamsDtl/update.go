package mstparamsdtl

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (m *mstParamsDtlRepository) UpdateParamsDtl(ctx context.Context, parmObj []schema.UpdateParamsDtl, sql sqlx.Sqlx) error {
	tx := m.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table("mst_params_dtl").Where(sql.Stmt, sql.Args...).Updates(&parmObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
