package usersgroupsetting

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
)

func (u *usersGroupSettingService) CreateUsersGroupSetting(ctx context.Context, req []request.CreateUsersGroupSettingRows) error {
	var parmObj []schema.CreateUsersGroupSetting
	for _, rec := range req {
		data := schema.CreateUsersGroupSetting{
			UserID:    rec.UserID,
			GroupID:   rec.GroupID,
			CreatedAt: rec.CreatedAt,
			DeletedAt: rec.DeletedAt,
			UpdatedAt: rec.UpdatedAt,
		}
		parmObj = append(parmObj, data)
	}

	if err := u.repos.CreateUsersGroupSetting(ctx, parmObj); err != nil {
		return err
	}

	return nil
}
