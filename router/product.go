package router

import (
	"SH-admin/app/api"
	"SH-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

//type ProductRouter struct {
//}
//
//func NewProductRouter() *ProductRouter {
//	ins := ProductRouter{}
//	return &ins
//}

func InitProductRouter(g *gin.RouterGroup) {
	api := api.NewProductApi()
	pg := g.Group("/product").Use(middleware.DbLogHandler())

	{
		pg.GET(":id", api.GetByIdApi)                  //http://localhost:5001/api/product/44
		pg.GET("GetByCodeApi/:code", api.GetByCodeApi) //http://localhost:5001/api/product/GetByCode/005
		pg.GET("", api.FindWithPagerApi)               //http://localhost:5001/api/product?pageNum=1&pageSize=2&code=024&unit=只
		pg.PUT("", api.InsertApi)
	}
}
