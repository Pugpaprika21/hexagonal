package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"

	"gorm.io/gorm"
)

type ISysLeftBarMenusRepository interface {
	GetMainManus(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetNewMainManus, error)
}

type sysLeftBarMenusRepository struct {
	db *gorm.DB
}

func NewSysLeftBarMenusRepository(db *gorm.DB) ISysLeftBarMenusRepository {
	return &sysLeftBarMenusRepository{db: db}
}
