package router

import (
	"SH-admin/api"
	"SH-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCustomerRouter(g *gin.RouterGroup) {
	cg := g.Group("customer").Use(middleware.DbLogHandler())
	{
		cg.GET(":id", api.NewCustomerApi().GetByIdApi)
		cg.GET("GetByEmailApi/:email", api.NewCustomerApi().GetByEmailApi)
		cg.GET("", api.NewCustomerApi().FindWithPagerApi)
		cg.PUT("", api.NewCustomerApi().UpdateApi)
	}
}
