package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *appRouter) v1() {
	api := r.server.Group("/api")

	v1 := api.Group("/v1", r.jwtx.Validate())
	{
		users := v1.Group("/users")
		{
			users.POST("/get", r.handler.Users.GetUsers)
		}
	}

	v1.GET("/heat_check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Is OK!")
	})
}
