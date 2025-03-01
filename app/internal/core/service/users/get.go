package users

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (u *usersService) GetUsers(ctx context.Context, req request.GetUsers) ([]response.GetUsers, error) {
	var sql sqlx.Sqlx

	sql.Stmt = `
        select 
            row_number() over (order by id desc) as row_num, 
            id, 
            username, 
            password, 
            first_name, 
            last_name, 
            date_of_birth, 
            address, 
            phone_number, 
            password_hash, 
            concat(first_name, ' ', last_name) as full_name
        from 
            users
    `

	if req.ID != nil && *req.ID != 0 {
		sql.WhereClause = append(sql.WhereClause, "id = ?")
		sql.Args = append(sql.Args, req.ID)
	}

	if req.Username != nil && *req.Password != "" {
		sql.WhereClause = append(sql.WhereClause, "username = ?")
		sql.Args = append(sql.Args, req.Username)
	}

	if req.Password != nil && *req.Password != "" {
		sql.WhereClause = append(sql.WhereClause, "password = ?")
		sql.Args = append(sql.Args, req.Password)
	}

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " where " + strings.Join(sql.WhereClause, " and ")
	}

	sql.Stmt += " order by id desc"

	if req.LazyLoad.Limit > 0 {
		sql.Stmt += " limit ?"
		sql.Args = append(sql.Args, req.LazyLoad.Limit)
	}

	if req.LazyLoad.Offset > 0 {
		sql.Stmt += " offset ?"
		sql.Args = append(sql.Args, req.LazyLoad.Offset)
	}

	rows, err := u.repos.GetUsers(ctx, sql)
	if err != nil {
		return nil, err
	}

	resp := make([]response.GetUsers, len(rows))
	for i, rec := range rows {
		resp[i] = response.GetUsers{
			RowNum:       rec.RowNum.Int64,
			ID:           int(rec.ID.Int64),
			Address:      rec.Address.String,
			Email:        rec.Email.String,
			FirstName:    rec.FirstName.String,
			LastName:     rec.LastName.String,
			Password:     rec.Password.String,
			FullName:     rec.FullName.String,
			PasswordHash: rec.PasswordHash.String,
			PhoneNumber:  rec.PhoneNumber.String,
			Username:     rec.Username.String,
			DateOfBirth:  rec.DateOfBirth.String,
		}
	}

	return resp, nil
}
