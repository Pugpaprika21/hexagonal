package users

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (u *usersService) CreateUsers(ctx context.Context, req []request.CreateUsersRows) error {
	parmObj := make([]schema.CreateUsers, len(req))
	for i, rec := range req {
		parmObj[i] = schema.CreateUsers{
			Address:      rec.Address,
			CreatedAt:    rec.CreatedAt,
			DateOfBirth:  rec.DateOfBirth,
			DeletedAt:    rec.DeletedAt,
			Email:        rec.Email,
			FirstName:    rec.FirstName,
			LastName:     rec.LastName,
			Password:     rec.Password,
			PasswordHash: rec.PasswordHash,
			PhoneNumber:  rec.PhoneNumber,
			UpdatedAt:    rec.UpdatedAt,
			Username:     sqlx.Nil(rec.Username),
		}
	}

	if err := u.repos.CreateUsers(ctx, parmObj); err != nil {
		return err
	}

	return nil
}
