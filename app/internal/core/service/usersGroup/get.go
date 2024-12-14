package usersgroup

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/respone"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersGroupService) GetUsersGroup(ctx context.Context, req request.GetUsersGroup) ([]respone.GetUsersGroup, error) {
	var sql sqlx.Sqlx
	var resp []respone.GetUsersGroup

	sql.Stmt = "select * from users_group"
	if req.ID != 0 {
		sql.WhereClause = append(sql.WhereClause, "id = ?")
		sql.Args = append(sql.Args, req.ID)
	}

	if req.GroupCode != "" {
		sql.WhereClause = append(sql.WhereClause, "group_code = ?")
		sql.Args = append(sql.Args, req.GroupCode)
	}

	if req.GroupName != "" {
		sql.WhereClause = append(sql.WhereClause, "group_name = ?")
		sql.Args = append(sql.Args, req.GroupName)
	}

	if req.GroupDescription != "" {
		sql.WhereClause = append(sql.WhereClause, "group_description = ?")
		sql.Args = append(sql.Args, req.GroupDescription)
	}

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	rows, err := u.repos.GetUsersGroup(ctx, sql)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		for _, rec := range rows {
			data := respone.GetUsersGroup{
				ID:               int(rec.ID.Int64),
				GroupCode:        rec.GroupCode.String,
				GroupName:        rec.GroupName.String,
				GroupDescription: rec.GroupDescription.String,
				UpdatedAt:        rec.UpdatedAt.String,
				CreatedAt:        rec.CreatedAt.String,
				DeletedAt:        rec.DeletedAt.String,
			}
			resp = append(resp, data)
		}
	}

	return resp, nil
}
