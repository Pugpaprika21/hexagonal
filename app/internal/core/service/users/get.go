package users

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/respone"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersSevice) GetUsers(ctx context.Context, req request.GetUsers) ([]respone.GetUsers, error) {
	var sql sqlx.Sqlx
	var resp []respone.GetUsers

	sql.Stmt = "select * from users"
	if req.ID != 0 {
		sql.WhereClause = append(sql.WhereClause, "id = ?")
		sql.Args = append(sql.Args, req.ID)
	}

	if req.Username != "" {
		sql.WhereClause = append(sql.WhereClause, "username = ?")
		sql.Args = append(sql.Args, req.Username)
	}

	if req.Password != "" {
		sql.WhereClause = append(sql.WhereClause, "password = ?")
		sql.Args = append(sql.Args, req.Password)
	}

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	sql.Stmt += " order by id desc"

	rows, err := u.repos.GetUsers(ctx, sql)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		for _, rec := range rows {
			data := respone.GetUsers{
				ID:           int(rec.ID.Int64),
				Address:      rec.Address.String,
				Email:        rec.Email.String,
				FirstName:    rec.FirstName.String,
				LastName:     rec.LastName.String,
				Password:     rec.Password.String,
				PasswordHash: rec.PasswordHash.String,
				PhoneNumber:  rec.PhoneNumber.String,
				UpdatedAt:    rec.UpdatedAt.String,
				Username:     rec.Username.String,
				CreatedAt:    rec.CreatedAt.String,
				DateOfBirth:  rec.DateOfBirth.String,
				DeletedAt:    rec.DeletedAt.String,
			}
			resp = append(resp, data)
		}
	}

	return resp, nil
}
