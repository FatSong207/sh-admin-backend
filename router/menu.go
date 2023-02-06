package router

import (
	"SH-admin/api"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(g *gin.RouterGroup) {
	pg := g.Group("menu")
	{
		pg.GET("/tree", api.NewMenuApi().GetMenuTreeApi)
		pg.GET("", api.NewMenuApi().GetAllMenuTreeApi)
		pg.GET("/cascader", api.NewMenuApi().GetAllMenuTreeCasApi)
		pg.POST("", api.NewMenuApi().InsertApi)
		pg.PUT("", api.NewMenuApi().UpdateApi)
		pg.GET(":id", api.NewMenuApi().GetByIdApi)
	}
}
