package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	sysleftbarmenus "pugpaprika/app/internal/adapter/repository/sysLeftbarMenus"
)

type ISysLeftBarMenusService interface {
	GetMainMenus(ctx context.Context, req request.GetNewMainManus) ([]response.GetNewMainManus, error)
	GetAllMenus(ctx context.Context, req request.GetAllMenus) ([]response.GetAllMenus, error)
}

type sysLeftBarMenusService struct {
	repos sysleftbarmenus.ISysLeftBarMenusRepository
}

func NewSysLeftBarMenusService(repos sysleftbarmenus.ISysLeftBarMenusRepository) ISysLeftBarMenusService {
	return &sysLeftBarMenusService{repos: repos}
}
