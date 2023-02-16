package router

import (
	"SH-admin/api"
	"SH-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLogRouter(g *gin.RouterGroup) {
	lg := g.Group("/logs").Use(middleware.DbLogHandler()) //若分組名為log則會被adblock擋掉：https://stackoverflow.com/questions/23341765/getting-neterr-blocked-by-client-error-on-some-ajax-calls
	{
		lg.GET("/:id", api.NewLogApi().GetByIdApi)
	}
	lgWithoutDbLog := g.Group("/logs")
	{
		lgWithoutDbLog.GET("", api.NewLogApi().FindWithPagerApi)
		lgWithoutDbLog.GET("/loginlogs", api.NewLogApi().FindLoginlogWithPagerApi)
	}
}
