package mstparamshdr

import (
	"context"
	"pugpaprika/app/internal/core/schema"
)

func (m *mstParamsHdrRepository) CreateParamsHdr(ctx context.Context, parmObj []schema.CreateParamsHdr) error {
	tx := m.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table("mst_params_hdr").Create(&parmObj).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
