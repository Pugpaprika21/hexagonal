package users

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	"pugpaprika/app/internal/adapter/repository/users"
)

type IUsersSevice interface {
	GetUsers(ctx context.Context, req request.GetUsers) ([]response.GetUsers, error)
	CreateUsers(ctx context.Context, req []request.CreateUsersRows) error
	UpdateUsers(ctx context.Context, req []request.UpdateUsersRows) error
	DeleteUsers(ctx context.Context, req request.DeleteUsers) error
}

type usersService struct {
	repos users.IUsersRepository
}

func NewUsersService(repos users.IUsersRepository) IUsersSevice {
	return &usersService{repos: repos}
}
