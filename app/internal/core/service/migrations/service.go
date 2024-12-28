package migrations

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/adapter/repository/migrations"
)

type IMigrationsService interface {
	GetGoStructProcedure(ctx context.Context, parmObj request.GetGoStructProcedure) (string, error)
	GetAllGoStructProcedure(ctx context.Context, parmObj request.GetAllGoStructProcedure) ([]string, error)
	CreateGoStructProcedure(ctx context.Context) error
}

type migrationsService struct {
	repos migrations.IMigrationsRepository
}

func NewMigrationsService(repos migrations.IMigrationsRepository) IMigrationsService {
	return &migrationsService{repos: repos}
}
