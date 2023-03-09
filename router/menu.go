package router

import (
	"SH-admin/app/api"
	"SH-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(g *gin.RouterGroup) {
	api := api.NewMenuApi()
	mg := g.Group("menu").Use(middleware.DbLogHandler())
	{
		mg.POST("", api.InsertApi)
		mg.PUT("", api.UpdateApi)
		mg.GET(":id", api.GetByIdApi)
		mg.DELETE(":ids", api.DeleteApi)
	}
	mWithoutDblog := g.Group("menu")
	{
		mWithoutDblog.GET("/tree", api.GetMenuTreeApi)
		mWithoutDblog.GET("", api.GetAllMenuTreeApi)
		mWithoutDblog.GET("/cascader", api.GetAllMenuTreeCasApi)
	}
}
