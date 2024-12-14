package usersgroup

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
)

func (u *usersGroupService) CreateUsersGroup(ctx context.Context, req []request.CreateUsersGroupRows) error {
	var parmObj []schema.CreateUsersGroup
	for _, rec := range req {
		data := schema.CreateUsersGroup{
			GroupCode:        rec.GroupCode,
			GroupName:        rec.GroupName,
			GroupDescription: rec.GroupDescription,
			CreatedAt:        rec.CreatedAt,
			DeletedAt:        rec.DeletedAt,
			UpdatedAt:        rec.UpdatedAt,
		}
		parmObj = append(parmObj, data)
	}

	if err := u.repos.CreateUsersGroup(ctx, parmObj); err != nil {
		return err
	}

	return nil
}
