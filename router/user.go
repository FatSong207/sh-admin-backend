package router

import (
	"SH-admin/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.RouterGroup) {
	//ug := g.Group("/user").Use(middleware.DbLogHandler())
	//{
	//
	//}
	ugWithoutDbLog := g.Group("/user")
	{
		ugWithoutDbLog.GET("/info", api.NewUserApi().GetUserInfoApi)
	}
}
