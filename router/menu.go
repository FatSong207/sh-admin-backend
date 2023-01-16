package router

import (
	"SH-admin/api"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(g *gin.RouterGroup) {
	pg := g.Group("/menu")
	{
		pg.GET("/tree", api.NewMenuApi().GetMenuTree)
		pg.GET("", api.NewMenuApi().GetAllMenuTree)
		pg.PUT("", api.NewMenuApi().BaseApi.Insert)
	}
}
