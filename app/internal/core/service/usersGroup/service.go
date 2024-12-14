package usersgroup

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	usersgroup "pugpaprika/app/internal/adapter/repository/usersGroup"
)

type IUsersGroupService interface {
	GetUsersGroup(ctx context.Context, req request.GetUsersGroup) ([]response.GetUsersGroup, error)
	CreateUsersGroup(ctx context.Context, req []request.CreateUsersGroupRows) error
	UpdateUsersGroup(ctx context.Context, req []request.UpdateUsersGroupRows) error
	DeleteUsersGroup(ctx context.Context, req request.DeleteUsersGroup) error
}

type usersGroupService struct {
	repos usersgroup.IUsersGroupRepository
}

func NewUsersGroupService(repos usersgroup.IUsersGroupRepository) IUsersGroupService {
	return &usersGroupService{repos: repos}
}
