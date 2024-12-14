package usersgroupsetting

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersGroupSettingService) GetUsersGroupSetting(ctx context.Context, req request.GetUsersGroupSetting) ([]response.GetUsersGroupSetting, error) {
	var sql sqlx.Sqlx
	var resp []response.GetUsersGroupSetting

	sql.Stmt = "select * from users_group_setting"
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

	rows, err := u.repos.GetUsersGroupSetting(ctx, sql)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		for _, rec := range rows {
			data := response.GetUsersGroupSetting{
				ID:        int(rec.ID.Int64),
				UserID:    int(rec.UserID.Int64),
				GroupID:   int(rec.GroupID.Int64),
				UpdatedAt: rec.UpdatedAt.String,
				CreatedAt: rec.CreatedAt.String,
				DeletedAt: rec.DeletedAt.String,
			}
			resp = append(resp, data)
		}
	}

	return resp, nil
}

func (u *usersGroupSettingService) GetUsersGroupSettingByUserID(ctx context.Context, req request.GetUsersGroupSettingByUserID) ([]response.GetUsersGroupSettingByUserID, error) {
	var sql sqlx.Sqlx
	var resp []response.GetUsersGroupSettingByUserID

	sql.Stmt = `
		select distinct ug.id, ug.group_name, case when ugs.user_id is not null then true else false end as has_in_group
		from users_group as ug
		left join (select group_id, user_id from users_group_setting where user_id = ?) as ugs on ug.id = ugs.group_id
	`
	if req.ID != 0 {
		sql.WhereClause = append(sql.WhereClause, "ug.id = ?")
		sql.Args = append(sql.Args, req.ID)
	}

	if req.UserID != 0 {
		sql.Args = append(sql.Args, req.UserID)
	} else {
		sql.Args = append(sql.Args, 0)
	}

	if req.GroupName != "" {
		sql.WhereClause = append(sql.WhereClause, "ug.group_name = ?")
		sql.Args = append(sql.Args, req.GroupName)
	}

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	sql.Stmt += " order by ug.created_at desc;"

	rows, err := u.repos.GetUsersGroupSettingByUserID(ctx, sql)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		for _, rec := range rows {
			data := response.GetUsersGroupSettingByUserID{
				ID:         int(rec.ID.Int64),
				GroupName:  rec.GroupName.String,
				HasInGroup: int(rec.HasInGroup.Int64),
				UpdatedAt:  rec.UpdatedAt.String,
				CreatedAt:  rec.CreatedAt.String,
				DeletedAt:  rec.DeletedAt.String,
			}
			resp = append(resp, data)
		}

	}

	return resp, nil
}
