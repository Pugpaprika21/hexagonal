package usersgroup

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (u *usersGroupRepository) GetUsersGroup(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetUsersGroup, error) {
	var rows []schema.GetUsersGroup
	result := u.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}
	return rows, nil
}
