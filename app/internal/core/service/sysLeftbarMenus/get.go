package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (s *sysLeftBarMenusService) GetMainMenus(ctx context.Context, req request.GetNewMainManus) ([]response.GetNewMainManus, error) {
	var sql sqlx.Sqlx
	var resp []response.GetNewMainManus

	sql.Stmt = `
	select 
		p.id as menu_id,
		p.name as menu_name,
		p.url as menu_url,
		p.icon as menu_icon,
		p.tooltip as menu_tooltip,
		p.position as menu_position,
		p.is_active as menu_is_active,
		s.id as sub_menu_id,
		s.name as sub_menu_name,
		s.url as sub_menu_url,
		s.icon as sub_menu_icon,
		s.tooltip as sub_menu_tooltip,
		s.position as sub_menu_position,
		s.is_active as sub_menu_is_active,
		p.user_id,
    	p.role_id
	from 
		sys_leftbar_menus p
	left join 
		sys_leftbar_menus s on p.id = s.parent_id 
	`
	if req.MenuID != nil && *req.MenuID != 0 {
		sql.WhereClause = append(sql.WhereClause, "p.menu_id = ?")
		sql.Args = append(sql.Args, req.MenuID)
	}

	if req.UserID != nil && *req.UserID != 0 {
		sql.WhereClause = append(sql.WhereClause, "p.user_id = ?")
		sql.Args = append(sql.Args, req.UserID)
	}

	if req.RoleID != nil && *req.RoleID != 0 {
		sql.WhereClause = append(sql.WhereClause, "p.role_id = ?")
		sql.Args = append(sql.Args, req.RoleID)
	}

	sql.WhereClause = append(sql.WhereClause, "p.is_active = ?")
	sql.Args = append(sql.Args, true)

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	sql.Stmt += ` order by p.position asc, s.position asc;`

	rows, err := s.repos.GetMainMenus(ctx, sql)
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

func (s *sysLeftBarMenusService) GetAllMenus(ctx context.Context, req request.GetAllMenus) ([]response.GetAllMenus, error) {
	var sql sqlx.Sqlx
	var resp []response.GetAllMenus

	sql.Stmt = `select id, user_id, role_id, name, url, icon, is_active from sys_leftbar_menus `

	if req.UserID != nil && *req.UserID != 0 {
		sql.WhereClause = append(sql.WhereClause, "user_id = ?")
		sql.Args = append(sql.Args, req.UserID)
	}

	if req.RoleID != nil && *req.RoleID != 0 {
		sql.WhereClause = append(sql.WhereClause, "role_id = ?")
		sql.Args = append(sql.Args, req.RoleID)
	}

	sql.WhereClause = append(sql.WhereClause, "is_active = ?")
	sql.Args = append(sql.Args, true)

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	sql.Stmt += ` order by id asc;`

	rows, err := s.repos.GetParentMenus(ctx, sql)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		for _, rec := range rows {
			sql.Stmt = `select name, url, icon, parent_id, is_active from sys_leftbar_menus where parent_id = ? and is_active = true`
			sql.Args = append(sql.Args[:0], rec.ID)
			childRows, err := s.repos.GetChildMenus(ctx, sql)
			if err != nil {
				return nil, err
			}

			data := response.GetAllMenus{
				ID:       rec.ID.Int32,
				UserID:   rec.UserID.Int32,
				RoleID:   rec.RoleID.Int32,
				Name:     rec.Name.String,
				Url:      rec.Url.String,
				Icon:     rec.Icon.String,
				IsActive: rec.IsActive.Bool,
			}

			if len(childRows) > 0 {
				var subMenus []response.GetAllMenus
				for _, subRec := range childRows {
					subMenu := response.GetAllMenus{
						Name:     subRec.Name.String,
						Url:      subRec.Url.String,
						Icon:     subRec.Icon.String,
						ParentID: subRec.ParentID.Int32,
						IsActive: subRec.IsActive.Bool,
					}
					subMenus = append(subMenus, subMenu)
				}
				data.SubMenus = subMenus
			}

			resp = append(resp, data)
		}
	}

	return resp, nil
}
