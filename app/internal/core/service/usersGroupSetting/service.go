package usersgroupsetting

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/respone"
	usersgroupsetting "pugpaprika/app/internal/adapter/repository/usersGroupSetting"
)

type IUsersGroupSettingService interface {
	GetUsersGroupSetting(ctx context.Context, req request.GetUsersGroupSetting) ([]respone.GetUsersGroupSetting, error)
	GetUsersGroupSettingByUserID(ctx context.Context, req request.GetUsersGroupSettingByUserID) ([]respone.GetUsersGroupSettingByUserID, error)
	CreateUsersGroupSetting(ctx context.Context, req []request.CreateUsersGroupSettingRows) error
	UpdateUsersGroupSetting(ctx context.Context, req []request.UpdateUsersGroupSettingRows) error
	DeleteUsersSettingGroup(ctx context.Context, req request.DeleteUsersSettingGroup) error
}

type usersGroupSettingService struct {
	repos usersgroupsetting.IUsersGroupSettingRepository
}

func NewUsersGroupSettingService(repos usersgroupsetting.IUsersGroupSettingRepository) IUsersGroupSettingService {
	return &usersGroupSettingService{repos: repos}
}
