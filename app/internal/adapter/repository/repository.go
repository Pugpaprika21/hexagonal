package repository

import (
	"pugpaprika/app/internal/adapter/repository/users"

	"gorm.io/gorm"
)

type Repository struct {
	Users users.IUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users: users.NewUserRepository(db),
	}
}
