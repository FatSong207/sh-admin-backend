package router

import (
	"SH-admin/app/api"
	"SH-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitCustomerRouter(g *gin.RouterGroup) {
	api := api.NewCustomerApi()
	cg := g.Group("customer").Use(middleware.DbLogHandler())
	{
		cg.GET(":id", api.GetByIdApi)
		cg.GET("GetByEmail/:email", api.GetByEmailApi)
		cg.PUT("", api.UpdateApi)
	}
	cgWithoutDbLog := g.Group("customer")
	{
		cgWithoutDbLog.GET("", api.FindWithPagerApi)
	}
}
