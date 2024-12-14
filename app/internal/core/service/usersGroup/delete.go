package usersgroup

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersGroupService) DeleteUsersGroup(ctx context.Context, req request.DeleteUsersGroup) error {
	var sql sqlx.Sqlx

	sql.Stmt = "delete from users_group "

	if req.ID != 0 {
		sql.WhereClause = append(sql.WhereClause, "id = ?")
		sql.Args = append(sql.Args, req.ID)
	}

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	if err := u.repos.DeleteUsersGroup(ctx, sql); err != nil {
		return err
	}

	return nil
}
