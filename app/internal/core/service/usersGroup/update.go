package usersgroup

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersGroupService) UpdateUsersGroup(ctx context.Context, req []request.UpdateUsersGroupRows) error {
	var sql sqlx.Sqlx
	var parmObj []schema.UpdateUsersGroup

	for _, rec := range req {
		data := schema.UpdateUsersGroup{
			GroupCode:        rec.GroupCode,
			GroupName:        rec.GroupName,
			GroupDescription: rec.GroupDescription,
			CreatedAt:        rec.CreatedAt,
			DeletedAt:        rec.DeletedAt,
			UpdatedAt:        rec.UpdatedAt,
		}

		parmObj = append(parmObj, data)

		if *rec.ID != 0 {
			sql.WhereClause = append(sql.WhereClause, "id = ?")
			sql.Args = append(sql.Args, rec.ID)
		}

		if len(sql.WhereClause) > 0 {
			sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
		}

		if err := u.repos.UpdateUsersGroup(ctx, parmObj, sql); err != nil {
			return err
		}
	}

	return nil
}
