package router

import (
	"SH-admin/api"
	"SH-admin/middleware"
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
	pg := g.Group("/product").Use(middleware.DbLogHandler()).Use(middleware.AuthorizeHandler())

	{
		pg.GET(":id", api.NewProductApi().GetByIdApi)                  //http://localhost:5001/api/product/44
		pg.GET("GetByCodeApi/:code", api.NewProductApi().GetByCodeApi) //http://localhost:5001/api/product/GetByCode/005
		pg.GET("", api.NewProductApi().FindWithPagerApi)               //http://localhost:5001/api/product?pageNum=1&pageSize=2&code=024&unit=只
		pg.PUT("", api.NewProductApi().InsertApi)
	}
}
