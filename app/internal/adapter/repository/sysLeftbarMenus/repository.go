package sysleftbarmenus

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"

	"gorm.io/gorm"
)

type ISysLeftBarMenusRepository interface {
	GetMainMenus(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetNewMainManus, error)
	GetParentMenus(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetParentMenus, error)
	GetChildMenus(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetChildMenus, error)
}

type sysLeftBarMenusRepository struct {
	db *gorm.DB
}

func NewSysLeftBarMenusRepository(db *gorm.DB) ISysLeftBarMenusRepository {
	return &sysLeftBarMenusRepository{db: db}
}
