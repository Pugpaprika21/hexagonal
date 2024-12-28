package router

import (
	"os"
	"pugpaprika/app/internal/adapter/handler"
	"pugpaprika/app/pkg/jwtx"

	"github.com/labstack/echo/v4"
)

type appRouter struct {
	server  *echo.Echo
	handler *handler.Handler
	jwtx    jwtx.IJwtx
}

func NewRouter(handler *handler.Handler) *appRouter {
	server := echo.New()
	return &appRouter{
		server:  server,
		handler: handler,
		jwtx:    jwtx.New(),
	}
}

func (r *appRouter) Start() {
	r.register()
	r.server.Logger.Fatal(r.server.Start(":" + os.Getenv("PORT")))
}

func (r *appRouter) register() {
	r.v1()
	r.v2()
}
