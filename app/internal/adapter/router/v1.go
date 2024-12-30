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
			users.POST("/getUsers", r.handler.Users.GetUsers)
			users.POST("/createUsers", r.handler.Users.CreateUsers)
			users.POST("/updateUsers", r.handler.Users.UpdateUsers)
			users.POST("/deleteUsers", r.handler.Users.DeleteUsers)
		}

		usersGroup := v1.Group("/usersGroup")
		{
			usersGroup.POST("/getUsersGroup", r.handler.UsersGroup.GetUsersGroup)
			usersGroup.POST("/createUsersGroup", r.handler.UsersGroup.CreateUsersGroup)
			usersGroup.POST("/updateUsersGroup", r.handler.UsersGroup.UpdateUsersGroup)
			usersGroup.POST("/deleteUsersGroup", r.handler.UsersGroup.DeleteUsersGroup)
		}

		usersGroupSetting := v1.Group("/usersGroupSetting")
		{
			usersGroupSetting.POST("/getUsersGroupSetting", r.handler.UsersGroupSetting.GetUsersGroupSetting)
			usersGroupSetting.POST("/getUsersGroupSettingByUserID", r.handler.UsersGroupSetting.GetUsersGroupSettingByUserID)
			usersGroupSetting.POST("/createUsersGroupSetting", r.handler.UsersGroupSetting.CreateUsersGroupSetting)
			usersGroupSetting.POST("/updateUsersGroupSetting", r.handler.UsersGroupSetting.UpdateUsersGroupSetting)
			usersGroupSetting.POST("/deleteUsersSettingGroup", r.handler.UsersGroupSetting.DeleteUsersSettingGroup)
		}

		migrations := v1.Group("/migrations")
		{
			migrations.POST("/getGoStructProcedure", r.handler.Migrations.GetGoStructProcedure)
			migrations.POST("/getAllGoStructProcedure", r.handler.Migrations.GetAllGoStructProcedure)
			migrations.POST("/createGoStructProcedure", r.handler.Migrations.CreateGoStructProcedure)
			migrations.POST("/deleteAllGoStructProcedure", r.handler.Migrations.DeleteAllGoStructProcedure)
		}

		sysLeftbarMenus := v1.Group("/sysLeftbarMenus")
		{
			sysLeftbarMenus.GET("/getMainManus", r.handler.SysLeftBarMenus.GetMainMenus)
			sysLeftbarMenus.GET("/getAllMenus", r.handler.SysLeftBarMenus.GetAllMenus)
		}
	}

	v1.GET("/heat_check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Is OK!")
	})
}
