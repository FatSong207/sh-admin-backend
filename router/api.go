package router

import (
	"SH-admin/app/api"
	"SH-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(g *gin.RouterGroup) {
	api := api.NewApiApi()
	ag := g.Group("/api").Use(middleware.DbLogHandler())
	{
		ag.POST("", api.InsertApi)
		ag.PUT("", api.UpdateApi)
		ag.GET(":id", api.GetByIdApi)
		ag.DELETE(":ids", api.DeleteApi)
	}
	agWithoutDbLog := g.Group("/api")
	{
		agWithoutDbLog.GET("", api.FindWithPagerApi)
		agWithoutDbLog.GET("tree", api.GetAllApiTree)
	}
}
