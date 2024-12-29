package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (s *sysLeftBarMenusService) GetMainManus(ctx context.Context, req request.GetNewMainManus) ([]response.GetNewMainManus, error) {
	var sql sqlx.Sqlx
	var resp []response.GetNewMainManus

	sql.Stmt = `
	select 
		m1.id as menu_id,
		m1.name as menu_name,
		m1.url as menu_url,
		m1.icon as menu_icon,
		m1.tooltip as menu_tooltip,
		m1.position as menu_position,
		m1.is_active as menu_is_active,
		m2.id as sub_menu_id,
		m2.name as sub_menu_name,
		m2.url as sub_menu_url,
		m2.icon as sub_menu_icon,
		m2.tooltip as sub_menu_tooltip,
		m2.position as sub_menu_position,
		m2.is_active as sub_menu_is_active,
		m1.user_id,
    	m1.role_id
	from 
		sys_leftbar_menus m1
	left join 
		sys_leftbar_menus m2 on m1.id = m2.parent_id 
	`
	if req.MenuID != nil && *req.MenuID != 0 {
		sql.WhereClause = append(sql.WhereClause, "menu_id = ?")
		sql.Args = append(sql.Args, req.MenuID)
	}

	if req.UserID != nil && *req.UserID != 0 {
		sql.WhereClause = append(sql.WhereClause, "user_id = ?")
		sql.Args = append(sql.Args, req.UserID)
	}

	if req.RoleID != nil && *req.RoleID != 0 {
		sql.WhereClause = append(sql.WhereClause, "role_id = ?")
		sql.Args = append(sql.Args, req.RoleID)
	}

	sql.WhereClause = append(sql.WhereClause, "m1.is_active = ?")
	sql.Args = append(sql.Args, true)

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	sql.Stmt += ` order by m1.position asc, m2.position asc;`

	rows, err := s.repos.GetMainManus(ctx, sql)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		for _, rec := range rows {
			data := response.GetNewMainManus{
				MenuID:          rec.MenuID.Int32,
				MenuName:        rec.MenuName.String,
				MenuURL:         rec.MenuURL.String,
				MenuIcon:        rec.MenuIcon.String,
				MenuTooltip:     rec.MenuTooltip.String,
				MenuPosition:    rec.MenuPosition.Int32,
				MenuIsActive:    rec.MenuIsActive.Bool,
				SubMenuID:       rec.SubMenuID.Int32,
				SubMenuName:     rec.SubMenuName.String,
				SubMenuURL:      rec.SubMenuURL.String,
				SubMenuIcon:     rec.SubMenuIcon.String,
				SubMenuTooltip:  rec.SubMenuTooltip.String,
				SubMenuPosition: rec.SubMenuPosition.Int32,
				SubMenuIsActive: rec.SubMenuIsActive.Bool,
				UserID:          rec.UserID.Int32,
				RoleID:          rec.RoleID.Int32,
			}
			resp = append(resp, data)
		}
	}

	return resp, nil
}
