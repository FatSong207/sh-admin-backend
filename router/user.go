package router

import (
	"SH-admin/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.RouterGroup) {
	ug := g.Group("/user")
	{
		ug.GET("/info", api.NewUserApi().GetUserInfoApi)
		//ug.GET("/verifycode", api.NewUserApi().GetVerifyCode)
	}
}
