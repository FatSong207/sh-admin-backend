package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/models"
	response "SH-admin/models/common"
	"github.com/gin-gonic/gin"
)

type ProductApi struct {
	*BaseApi[models.Product, models.ProductOutDto]
	IServices.IProductService
}

func NewProductApi() *ProductApi {
	ins := &ProductApi{
		NewBaseApi[models.Product, models.ProductOutDto](),
		Services.NewProductService(),
	}
	return ins
}

// GetById
// @Summary 根據Id獲取Product對應的OutDto
// @Description 根據Id獲取Product對應的OutDto
// @Tags ProductApi
// @Accept json
// @Param id path int true "id主鍵"
// @Success 200 {object} response.Response{}
// @Router /product/{id} [get]
func (p *ProductApi) GetById(ctx *gin.Context) {
	p.BaseApi.GetById(ctx)
}

// GetByCode
// @Summary 根據Code獲取Product
// @Description 根據Code獲取實體Product
// @Tags ProductApi
// @Accept json
// @Param code path string true "code"
// @Success 200 {object} response.Response{}
// @Router /product/GetByCode/{code} [get]
func (p *ProductApi) GetByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	byCode, err := p.IProductService.GetByCode(code)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, byCode, ctx)
}

// FindWithPager
// @Summary Product分頁列表
// @Description Product分頁列表
// @Tags ProductApi
// @Accept json
// @Param data query response.PageInfo false "頁碼,單頁大小"
// @Param T query models.Product false "Product條件"
// @Success 200 {object} response.Response{}
// @Router /products [get]
func (p *ProductApi) FindWithPager(ctx *gin.Context) {
	var param = models.NewSearchDto[models.Product]()
	err := ctx.ShouldBind(param)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	withPager, i, err := p.IProductService.FindWithPager(*param)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.PageResult(response.ErrCodeSuccess, withPager, i, ctx)
}
