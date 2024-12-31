package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (s *sysLeftBarMenusService) CreateMenus(ctx context.Context, req []request.CreateMenusRows) error {
	parmObj := make([]schema.CreateMenus, len(req))
	for i, rec := range req {
		parmObj[i] = schema.CreateMenus{
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
			CreatedAt:     sqlx.Nil(rec.CreatedAt),
			UpdatedAt:     sqlx.Nil(rec.UpdatedAt),
		}
	}

	if err := s.repos.CreateMenus(ctx, parmObj); err != nil {
		return err
	}

	return nil
}
