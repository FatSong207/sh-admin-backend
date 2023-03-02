package router

import (
	"SH-admin/api"
	"SH-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.RouterGroup) {
	ug := g.Group("/user").Use(middleware.DbLogHandler())
	{
		ug.GET("/:id", api.NewUserApi().GetByIdApi)
		ug.PUT("", api.NewUserApi().UpdateApi)
		ug.POST("/sendmail", api.NewUserApi().SendMailToUserApi)
	}
	ugWithoutDbLog := g.Group("/user")
	{
		ugWithoutDbLog.GET("", api.NewUserApi().FindWithPagerApi)
	}
}
