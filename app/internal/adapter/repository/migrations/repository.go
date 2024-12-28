package migrations

import (
	"context"
	"pugpaprika/app/pkg/sqlx"

	"gorm.io/gorm"
)

type IMigrationsRepository interface {
	GetGoStructProcedure(ctx context.Context, sql sqlx.Sqlx) ([]string, error)
	GetAllGoStructProcedure(ctx context.Context, sql sqlx.Sqlx) ([]string, error)
	CreateGoStructProcedure(ctx context.Context, sql sqlx.Sqlx) error
}

type migrationsRepository struct {
	db *gorm.DB
}

func NewMigrationsRepository(db *gorm.DB) IMigrationsRepository {
	return &migrationsRepository{db: db}
}
