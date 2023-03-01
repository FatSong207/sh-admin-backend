package router

import (
	"SH-admin/api"
	"SH-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(g *gin.RouterGroup) {
	cg := g.Group("/casbin").Use(middleware.DbLogHandler())
	{
		cg.PUT("", api.NewCasbinApi().UpdateCasbin)
	}
	cgWithoutDbLog := g.Group("/casbin")
	{
		cgWithoutDbLog.GET("/:roleid", api.NewCasbinApi().GetAccessApiPathByRoleId)
	}
}
