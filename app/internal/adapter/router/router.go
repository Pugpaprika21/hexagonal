package router

import (
	"os"

	"github.com/labstack/echo/v4"
)

type appRouter struct {
	server *echo.Echo
}

func NewRouter() *appRouter {
	server := echo.New()
	return &appRouter{server: server}
}

func (r *appRouter) register() {
	r.v1()
	r.v2()
}

func (r *appRouter) Start() {
	r.register()
	r.server.Logger.Fatal(r.server.Start(":" + os.Getenv("PORT")))
}
