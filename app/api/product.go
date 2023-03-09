package api

import (
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/services"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"github.com/gin-gonic/gin"
)

type ProductApi struct {
	*BaseApi[models.Product, models.ProductOutDto]
	iService interface_services.IProductService
}

func NewProductApi() *ProductApi {
	ins := &ProductApi{
		NewBaseApi[models.Product, models.ProductOutDto](),
		services.NewProductService(),
	}
	return ins
}

// GetByIdApi
// @Summary 根據Id獲取Product對應的OutDto
// @Description 根據Id獲取Product對應的OutDto
// @Tags ProductApi
// @Accept json
// @Param id path int true "id主鍵"
// @Success 200 {object} common.common{}
// @Router /product/{id} [get]
func (p *ProductApi) GetByIdApi(ctx *gin.Context) {
	p.BaseApi.GetByIdApi(ctx)
}

// GetByCodeApi
// @Summary 根據Code獲取Product
// @Description 根據Code獲取實體Product
// @Tags ProductApi
// @Accept json
// @Param code path string true "code"
// @Success 200 {object} common.common{}
// @Router /product/GetByCodeApi/{code} [get]
func (p *ProductApi) GetByCodeApi(ctx *gin.Context) {
	code := ctx.Param("code")
	byCode, err := p.iService.GetByCode(code)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, byCode, ctx)
}

// FindWithPagerApi
// @Summary Product分頁列表
// @Description Product分頁列表
// @Tags ProductApi
// @Accept json
// @Param data query common.PageInfo false "頁碼,單頁大小"
// @Param T query models.Product false "Product條件"
// @Success 200 {object} common.common{}
// @Router /products [get]
func (p *ProductApi) FindWithPagerApi(ctx *gin.Context) {
	var param = common.NewSearchDto[models.Product]()
	err := ctx.ShouldBind(param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	withPager, i, err := p.iService.FindWithPager(*param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.PageResult(common.ErrCodeSuccess, withPager, i, ctx)
}
