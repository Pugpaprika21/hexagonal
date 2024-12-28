package migrations

import (
	"context"
	"pugpaprika/app/pkg/sqlx"
)

func (m *migrationsRepository) DeleteAllGoStructProcedure(ctx context.Context, sql sqlx.Sqlx) error {
	tx := m.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	result := tx.Raw(sql.Stmt, sql.Args...)
	if result.Error != nil {
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
