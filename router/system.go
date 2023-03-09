package router

import (
	"SH-admin/app/api"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(g *gin.RouterGroup) {
	api := api.NewSystemApi()
	//sg := g.Group("/system")
	//{
	//
	//}
	sgWithoutDblog := g.Group("system")
	{
		sgWithoutDblog.GET("/serverinfo", api.GetServerInfoApi)
		sgWithoutDblog.GET("/dashboard", api.GetDashboardApi)
	}
}
