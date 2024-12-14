package usersgroupsetting

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersGroupSettingService) UpdateUsersGroupSetting(ctx context.Context, req []request.UpdateUsersGroupSettingRows) error {
	var sql sqlx.Sqlx
	var parmObj []schema.UpdateUsersGroupSetting

	for _, rec := range req {
		data := schema.UpdateUsersGroupSetting{
			UserID:    rec.UserID,
			GroupID:   rec.GroupID,
			CreatedAt: rec.CreatedAt,
			DeletedAt: rec.DeletedAt,
			UpdatedAt: rec.UpdatedAt,
		}

		parmObj = append(parmObj, data)

		if *rec.ID != 0 {
			sql.WhereClause = append(sql.WhereClause, "id = ?")
			sql.Args = append(sql.Args, rec.ID)
		}

		if *rec.UserID != 0 {
			sql.WhereClause = append(sql.WhereClause, "user_id = ?")
			sql.Args = append(sql.Args, rec.UserID)
		}

		if *rec.GroupID != 0 {
			sql.WhereClause = append(sql.WhereClause, "group_id = ?")
			sql.Args = append(sql.Args, rec.GroupID)
		}

		if len(sql.WhereClause) > 0 {
			sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
		}

		if err := u.repos.UpdateUsersGroupSetting(ctx, parmObj, sql); err != nil {
			return err
		}
	}

	return nil
}
