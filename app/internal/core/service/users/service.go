package users

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/respone"
	"pugpaprika/app/internal/adapter/repository/users"
)

type IUsersSevice interface {
	GetUsers(ctx context.Context, req request.GetUsers) ([]respone.GetUsers, error)
}

type usersSevice struct {
	users users.IUserRepository
}

func NewUsersSevice(users users.IUserRepository) IUsersSevice {
	return &usersSevice{users: users}
}
