package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (s *sysLeftBarMenusService) UpdateMenus(ctx context.Context, req []request.UpdateMenusRows) error {
	var sql sqlx.Sqlx
	parmObj := make([]schema.UpdateMenus, len(req))

	for i, rec := range req {
		parmObj[i] = schema.UpdateMenus{
			UserID:        sqlx.Nil(rec.UserID),
			RoleID:        sqlx.Nil(rec.RoleID),
			Name:          sqlx.Nil(rec.Name),
			NameEn:        sqlx.Nil(rec.NameEn),
			Url:           sqlx.Nil(rec.Url),
			Icon:          sqlx.Nil(rec.Icon),
			Tooltip:       sqlx.Nil(rec.Tooltip),
			ParentID:      sqlx.Nil(rec.ParentID),
			Position:      sqlx.Nil(rec.Position),
			IsActive:      sqlx.Nil(rec.IsActive),
			IsExternal:    sqlx.Nil(rec.IsExternal),
			PermissionKey: sqlx.Nil(rec.PermissionKey),
			UpdatedAt:     sqlx.Nil(rec.UpdatedAt),
		}

		if *rec.ID != 0 {
			sql.WhereClause = append(sql.WhereClause, "id = ?")
			sql.Args = append(sql.Args, rec.ID)
		}

		if *rec.UserID != 0 {
			sql.WhereClause = append(sql.WhereClause, "user_id = ?")
			sql.Args = append(sql.Args, rec.UserID)
		}

		if *rec.RoleID != 0 {
			sql.WhereClause = append(sql.WhereClause, "role_id = ?")
			sql.Args = append(sql.Args, rec.RoleID)
		}

		if len(sql.WhereClause) > 0 {
			sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
		}

		if err := s.repos.UpdateMenus(ctx, parmObj, sql); err != nil {
			return err
		}
	}

	return nil
}
