package usersgroup

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersGroupService) GetUsersGroup(ctx context.Context, req request.GetUsersGroup) ([]response.GetUsersGroup, error) {
	var sql sqlx.Sqlx

	sql.Stmt = "select id, group_code, group_name, group_description from users_group"
	if req.ID != nil && *req.ID != 0 {
		sql.WhereClause = append(sql.WhereClause, "id = ?")
		sql.Args = append(sql.Args, req.ID)
	}

	if req.GroupCode != nil && *req.GroupCode != "" {
		sql.WhereClause = append(sql.WhereClause, "group_code = ?")
		sql.Args = append(sql.Args, req.GroupCode)
	}

	if req.GroupName != nil && *req.GroupName != "" {
		sql.WhereClause = append(sql.WhereClause, "group_name = ?")
		sql.Args = append(sql.Args, req.GroupName)
	}

	if req.GroupDescription != nil && *req.GroupDescription != "" {
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

	resp := make([]response.GetUsersGroup, len(rows))
	for i, rec := range rows {
		resp[i] = response.GetUsersGroup{
			ID:               int(rec.ID.Int64),
			GroupCode:        rec.GroupCode.String,
			GroupName:        rec.GroupName.String,
			GroupDescription: rec.GroupDescription.String,
		}
	}

	return resp, nil
}
