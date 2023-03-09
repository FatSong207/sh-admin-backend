package router

import (
	"SH-admin/app/api"
	"SH-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.RouterGroup) {
	api := api.NewUserApi()
	ug := g.Group("/user").Use(middleware.DbLogHandler())
	{
		//ug.GET("/:id", api.GetByIdApi)
		//ug.GET("/:id", api.Test)
		ug.PUT("", api.UpdateApi)
		ug.POST("/sendmail", api.SendMailToUserApi)
	}
	ugWithoutDbLog := g.Group("/user")
	{
		ugWithoutDbLog.GET("", api.FindWithPagerApi)
	}
}
