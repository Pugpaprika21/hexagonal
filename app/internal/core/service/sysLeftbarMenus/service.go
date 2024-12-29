package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	sysleftbarmenus "pugpaprika/app/internal/adapter/repository/sysLeftbarMenus"
)

type ISysLeftBarMenusService interface {
	GetMainManus(ctx context.Context, req request.GetNewMainManus) ([]response.GetNewMainManus, error)
}

type sysLeftBarMenusService struct {
	repos sysleftbarmenus.ISysLeftBarMenusRepository
}

func NewSysLeftBarMenusService(repos sysleftbarmenus.ISysLeftBarMenusRepository) ISysLeftBarMenusService {
	return &sysLeftBarMenusService{repos: repos}
}
