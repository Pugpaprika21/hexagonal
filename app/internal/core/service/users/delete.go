package users

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersService) DeleteUsers(ctx context.Context, req request.DeleteUsers) error {
	var sql sqlx.Sqlx

	sql.Stmt = "delete from users "

	if req.ID != 0 {
		sql.WhereClause = append(sql.WhereClause, "id = ?")
		sql.Args = append(sql.Args, req.ID)
	}

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	if err := u.repos.DeleteUsers(ctx, sql); err != nil {
		return err
	}

	return nil
}
