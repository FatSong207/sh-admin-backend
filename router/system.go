package router

import (
	"SH-admin/api"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(g *gin.RouterGroup) {
	sg := g.Group("/system")
	{
		sg.GET("/serverinfo", api.NewSystemApi().GetServerInfo)
		//ug.GET("/verifycode", api.NewUserApi().GetVerifyCode)
	}
}
