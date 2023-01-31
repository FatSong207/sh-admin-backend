package router

import (
	"SH-admin/api"
	"github.com/gin-gonic/gin"
)

func InitCustomerRouter(g *gin.RouterGroup) {
	cg := g.Group("customer")
	{
		cg.GET(":id", api.NewCustomerApi().GetByIdApi)
		cg.GET("GetByEmailApi/:email", api.NewCustomerApi().GetByEmailApi)
		cg.GET("", api.NewCustomerApi().FindWithPagerApi)
	}
}
