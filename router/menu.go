package router

import (
	"SH-admin/api"
	"SH-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(g *gin.RouterGroup) {
	mg := g.Group("menu").Use(middleware.DbLogHandler()).Use(middleware.AuthorizeHandler())
	{
		mg.POST("", api.NewMenuApi().InsertApi)
		mg.PUT("", api.NewMenuApi().UpdateApi)
		mg.GET(":id", api.NewMenuApi().GetByIdApi)
		mg.DELETE(":ids", api.NewMenuApi().DeleteApi)
	}
	mWithoutDblog := g.Group("menu")
	{
		mWithoutDblog.GET("/tree", api.NewMenuApi().GetMenuTreeApi)
		mWithoutDblog.GET("", api.NewMenuApi().GetAllMenuTreeApi)
		mWithoutDblog.GET("/cascader", api.NewMenuApi().GetAllMenuTreeCasApi)
	}
}
