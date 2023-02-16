package router

import (
	"SH-admin/api"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(g *gin.RouterGroup) {
	//sg := g.Group("/system")
	//{
	//
	//}
	sgWithoutDblog := g.Group("system")
	{
		sgWithoutDblog.GET("/serverinfo", api.NewSystemApi().GetServerInfoApi)
		sgWithoutDblog.GET("/dashboard", api.NewSystemApi().GetDashboardApi)
	}
}
