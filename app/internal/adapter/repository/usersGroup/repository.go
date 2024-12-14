package usersgroup

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"

	"gorm.io/gorm"
)

type IUsersGroupRepository interface {
	GetUsersGroup(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetUsersGroup, error)
	CreateUsersGroup(ctx context.Context, parmObj []schema.CreateUsersGroup) error
	UpdateUsersGroup(ctx context.Context, parmObj []schema.UpdateUsersGroup, sql sqlx.Sqlx) error
	DeleteUsersGroup(ctx context.Context, sql sqlx.Sqlx) error
}

type usersGroupRepository struct {
	db *gorm.DB
}

func NewUsersGroupRepository(db *gorm.DB) IUsersGroupRepository {
	return &usersGroupRepository{db: db}
}
