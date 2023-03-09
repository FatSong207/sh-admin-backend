package router

import (
	"SH-admin/app/api"
	"SH-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitLogRouter(g *gin.RouterGroup) {
	api := api.NewLogApi()
	lg := g.Group("/logs").Use(middleware.DbLogHandler()) //若分組名為log則會被adblock擋掉：https://stackoverflow.com/questions/23341765/getting-neterr-blocked-by-client-error-on-some-ajax-calls
	{
		lg.GET("/:id", api.GetByIdApi)
	}
	lgWithoutDbLog := g.Group("/logs")
	{
		lgWithoutDbLog.GET("", api.FindWithPagerApi)
		lgWithoutDbLog.GET("/loginlogs", api.FindLoginlogWithPagerApi)
	}
}
