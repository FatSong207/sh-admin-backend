package router

import (
	"SH-admin/api"
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
	pg := g.Group("/product")

	{
		pg.GET(":id", api.NewProductApi().GetById)               //http://localhost:5001/api/product/44
		pg.GET("GetByCode/:code", api.NewProductApi().GetByCode) //http://localhost:5001/api/product/GetByCode/005
		pg.GET("", api.NewProductApi().FindWithPager)            //http://localhost:5001/api/product?pageNum=1&pageSize=2&code=024&unit=只
		pg.PUT("", api.NewProductApi().BaseApi.Insert)
	}
}
