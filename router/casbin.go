package router

import (
	"SH-admin/app/api"
	"SH-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(g *gin.RouterGroup) {
	api := api.NewCasbinApi()
	cg := g.Group("/casbin").Use(middleware.DbLogHandler())
	{
		cg.PUT("", api.UpdateCasbin)
	}
	cgWithoutDbLog := g.Group("/casbin")
	{
		cgWithoutDbLog.GET("/:roleid", api.GetAccessApiPathByRoleId)
	}
}
