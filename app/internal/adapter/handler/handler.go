package handler

import (
	"pugpaprika/app/internal/adapter/handler/users"
	"pugpaprika/app/internal/core/service"
)

type Handler struct {
	Users users.IUserHandler
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		Users: users.NewUserHandler(serv.Users),
	}
}
