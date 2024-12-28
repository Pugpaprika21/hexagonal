package migrations

import (
	"context"
	"pugpaprika/app/pkg/sqlx"
)

func (m *migrationsRepository) GetGoStructProcedure(ctx context.Context, sql sqlx.Sqlx) ([]string, error) {
	rows, err := m.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var genGoStruct string
		if err := rows.Scan(&genGoStruct); err != nil {
			return nil, err
		}
		result = append(result, genGoStruct)
	}

	return result, nil
}

func (m *migrationsRepository) GetAllGoStructProcedure(ctx context.Context, sql sqlx.Sqlx) ([]string, error) {
	rows, err := m.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var genGoStruct string
		if err := rows.Scan(&genGoStruct); err != nil {
			return nil, err
		}
		result = append(result, genGoStruct)
	}

	return result, nil
}
