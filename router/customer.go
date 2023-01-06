package router

import (
	"SH-admin/api"
	"github.com/gin-gonic/gin"
)

func InitCustomerRouter(g *gin.RouterGroup) {
	cg := g.Group("customer")
	{
		cg.GET(":id", api.NewCustomerApi().GetById)
		cg.GET("GetByEmail/:email", api.NewCustomerApi().GetByEmail)
		cg.GET("", api.NewCustomerApi().FindWithPager)
	}
}
