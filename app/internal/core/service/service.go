package service

import (
	"pugpaprika/app/internal/adapter/repository"
	"pugpaprika/app/internal/core/service/users"
)

type Service struct {
	Users users.IUsersSevice
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Users: users.NewUsersSevice(repos.Users),
	}
}
