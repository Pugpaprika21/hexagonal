package migrations

import (
	"context"
	"pugpaprika/app/pkg/sqlx"
)

func (m *migrationsRepository) CreateGoStructProcedure(ctx context.Context, sql sqlx.Sqlx) error {
	tx := m.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := m.db.Exec(sql.Stmt).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
