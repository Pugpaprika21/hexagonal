package usersgroupsetting

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
)

func (u *usersGroupSettingService) CreateUsersGroupSetting(ctx context.Context, req []request.CreateUsersGroupSettingRows) error {
	parmObj := make([]schema.CreateUsersGroupSetting, len(req))
	for i, rec := range req {
		parmObj[i] = schema.CreateUsersGroupSetting{
			UserID:    rec.UserID,
			GroupID:   rec.GroupID,
			CreatedAt: rec.CreatedAt,
			DeletedAt: rec.DeletedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}

	if err := u.repos.CreateUsersGroupSetting(ctx, parmObj); err != nil {
		return err
	}

	return nil
}
