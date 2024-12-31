package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (s *sysLeftBarMenusService) DeleteMenus(ctx context.Context, req request.DeleteMenus) error {
	var sql sqlx.Sqlx

	sql.Stmt = "delete from sys_leftbar_menus "

	if *req.ID != 0 {
		sql.WhereClause = append(sql.WhereClause, "id = ?")
		sql.Args = append(sql.Args, req.ID)
	}

	if *req.UserID != 0 {
		sql.WhereClause = append(sql.WhereClause, "user_id = ?")
		sql.Args = append(sql.Args, req.UserID)
	}

	if *req.RoleID != 0 {
		sql.WhereClause = append(sql.WhereClause, "role_id = ?")
		sql.Args = append(sql.Args, req.RoleID)
	}

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	if err := s.repos.DeleteMenus(ctx, sql); err != nil {
		return err
	}

	return nil
}
