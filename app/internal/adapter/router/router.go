package router

import (
	"os"
	"pugpaprika/app/internal/adapter/handler"

	"github.com/labstack/echo/v4"
)

type appRouter struct {
	server  *echo.Echo
	handler *handler.Handler
}

func NewRouter(handler *handler.Handler) *appRouter {
	server := echo.New()
	return &appRouter{server: server, handler: handler}
}

func (r *appRouter) register() {
	r.v1()
	r.v2()
}

func (r *appRouter) Start() {
	r.register()
	r.server.Logger.Fatal(r.server.Start(":" + os.Getenv("PORT")))
}
