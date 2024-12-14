package usersgroupsetting

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersGroupSettingService) DeleteUsersSettingGroup(ctx context.Context, req request.DeleteUsersSettingGroup) error {
	var sql sqlx.Sqlx

	sql.Stmt = "delete from users_group_setting "

	if req.ID != 0 {
		sql.WhereClause = append(sql.WhereClause, "id = ?")
		sql.Args = append(sql.Args, req.ID)
	}

	if req.UserID != 0 {
		sql.WhereClause = append(sql.WhereClause, "user_id = ?")
		sql.Args = append(sql.Args, req.UserID)
	}

	if req.GroupID != 0 {
		sql.WhereClause = append(sql.WhereClause, "group_id = ?")
		sql.Args = append(sql.Args, req.GroupID)
	}

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	if err := u.repos.DeleteUsersSettingGroup(ctx, sql); err != nil {
		return err
	}

	return nil
}
