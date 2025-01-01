package mstparamsdtl

import (
	"context"
	"pugpaprika/app/internal/core/schema"
)

func (m *mstParamsDtlRepository) CreateParamsDtl(ctx context.Context, parmObj []schema.CreateParamsDtl) error {
	tx := m.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table("mst_params_dtl").Create(&parmObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
