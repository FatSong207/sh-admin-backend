package router

import (
	"SH-admin/app/api"
	"SH-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitTaskJobRouter(g *gin.RouterGroup) {
	api := api.NewTaskJobApi()
	tjg := g.Group("/taskjob").Use(middleware.DbLogHandler())
	{
		tjg.GET("/:id", api.GetByIdApi)
		tjg.PUT("/status", api.ChangStatus)
		tjg.PUT("", api.UpdateApi)
	}
	tjgWithoutDbLog := g.Group("/taskjob")
	{
		tjgWithoutDbLog.GET("", api.FindWithPagerApi)
	}
}
