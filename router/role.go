package router

import (
	"SH-admin/api"
	"SH-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(g *gin.RouterGroup) {
	rg := g.Group("/role").Use(middleware.DbLogHandler())
	{
		rg.GET("/:id", api.NewRoleApi().GetByIdApi)
		rg.POST("", api.NewRoleApi().InsertApi)
		rg.PUT("", api.NewRoleApi().UpdateApi)
		rg.DELETE("/:ids", api.NewRoleApi().DeleteApi)
	}
	rgWithoutDbRole := g.Group("/role")
	{
		rgWithoutDbRole.GET("", api.NewRoleApi().FindWithPagerApi)
	}
}
