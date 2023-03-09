package router

import (
	"SH-admin/app/api"
	"SH-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(g *gin.RouterGroup) {
	api := api.NewRoleApi()
	rg := g.Group("/role").Use(middleware.DbLogHandler())
	{
		rg.GET("/:id", api.GetByIdApi)
		rg.POST("", api.InsertApi)
		rg.PUT("", api.UpdateApi)
		rg.DELETE("/:ids", api.DeleteApi)
	}
	rgWithoutDbRole := g.Group("/role")
	{
		rgWithoutDbRole.GET("", api.FindWithPagerApi)
		rgWithoutDbRole.GET("/all", api.GetAllRoleListForSelect)
	}
}
