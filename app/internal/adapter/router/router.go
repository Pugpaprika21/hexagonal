package router

import (
	"os"
	"pugpaprika/app/internal/adapter/handler"
	"pugpaprika/app/pkg/jwtx"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type appRouter struct {
	server  *echo.Echo
	handler *handler.Handler
	jwtx    jwtx.IJwtx
	dbTest  *gorm.DB
}

func NewRouter(handler *handler.Handler, dbTest *gorm.DB) *appRouter {
	server := echo.New()
	return &appRouter{
		server:  server,
		handler: handler,
		jwtx:    jwtx.New(),
		dbTest:  dbTest,
	}
}

func (r *appRouter) Start() {
	r.register()
	r.server.Logger.Fatal(r.server.Start(":" + os.Getenv("PORT")))
}

func (r *appRouter) register() {
	r.v1()
	r.v2()
	r.test()
}
