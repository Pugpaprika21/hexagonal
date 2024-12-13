package users

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (u *userRepository) GetUsers(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetUsers, error) {
	var rows []schema.GetUsers
	result := u.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}
