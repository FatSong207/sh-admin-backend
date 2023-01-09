package router

import (
	"SH-admin/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.RouterGroup) {
	pg := g.Group("/user")
	{
		pg.GET("/info", api.NewUserApi().GetUserInfo)
	}
}
