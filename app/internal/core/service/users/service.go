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
	repos users.IUserRepository
}

func NewUsersSevice(repos users.IUserRepository) IUsersSevice {
	return &usersSevice{repos: repos}
}
