package users

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"

	"gorm.io/gorm"
)

type IUsersRepository interface {
	GetUsers(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetUsers, error)
	CreateUsers(ctx context.Context, parmObj []schema.CreateUsers) error
	UpdateUsers(ctx context.Context, parmObj []schema.UpdateUsers, sql sqlx.Sqlx) error
	DeleteUsers(ctx context.Context, sql sqlx.Sqlx) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) IUsersRepository {
	return &userRepository{db: db}
}
