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
	CreateMenus(ctx context.Context, parmObj []schema.CreateMenus) error
	UpdateMenus(ctx context.Context, parmObj []schema.UpdateMenus, sql sqlx.Sqlx) error
	DeleteMenus(ctx context.Context, sql sqlx.Sqlx) error
}

type sysLeftBarMenusRepository struct {
	db *gorm.DB
}

func NewSysLeftBarMenusRepository(db *gorm.DB) ISysLeftBarMenusRepository {
	return &sysLeftBarMenusRepository{db: db}
}
