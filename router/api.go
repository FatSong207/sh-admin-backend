package router

import (
	"SH-admin/api"
	"SH-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(g *gin.RouterGroup) {
	ag := g.Group("/api").Use(middleware.DbLogHandler())
	{
		ag.POST("", api.NewApiApi().InsertApi)
		ag.PUT("", api.NewApiApi().UpdateApi)
		ag.GET(":id", api.NewApiApi().GetByIdApi)
		ag.DELETE(":ids", api.NewApiApi().DeleteApi)
	}
	agWithoutDbLog := g.Group("/api")
	{
		agWithoutDbLog.GET("", api.NewApiApi().FindWithPagerApi)
		agWithoutDbLog.GET("tree", api.NewApiApi().GetAllApiTree)
	}
}
