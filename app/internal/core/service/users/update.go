package users

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersSevice) UpdateUsers(ctx context.Context, req []request.UpdateUsersRows) error {
	var sql sqlx.Sqlx
	var parmObj []schema.UpdateUsers

	for _, rec := range req {
		data := schema.UpdateUsers{
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
			Username:     rec.Username,
		}

		parmObj = append(parmObj, data)

		if *rec.ID != 0 {
			sql.WhereClause = append(sql.WhereClause, "id = ?")
			sql.Args = append(sql.Args, rec.ID)
		}

		if len(sql.WhereClause) > 0 {
			sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
		}

		if err := u.repos.UpdateUsers(ctx, parmObj, sql); err != nil {
			return err
		}
	}

	return nil
}
