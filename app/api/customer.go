package api

import (
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/services"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"github.com/gin-gonic/gin"
)

type CustomerApi struct {
	*BaseApi[models.Customer, models.CustomerOutDto]
	iService interface_services.ICustomerService
}

func NewCustomerApi() *CustomerApi {
	ins := &CustomerApi{
		NewBaseApi[models.Customer, models.CustomerOutDto](),
		services.NewCustomerService(),
	}
	return ins
}

// GetByEmailApi
// @Summary 根據Email獲取Customer
// @Description 分頁列表
// @Tags CustomerApi
// @Accept json
// @Param email path string true "EMail"
// @Success 200 {object} common.common{}
// @Router /customer/GetByEmailApi/{email} [get]
func (c *CustomerApi) GetByEmailApi(ctx *gin.Context) {
	e := ctx.Param("email")
	t, err := c.iService.GetByEmail(e)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, t, ctx)
}

// GetByIdApi @Summary 根據Id獲取Customer對應的OutDto
// @Description 根據Id獲取Customer對應的OutDto
// @Tags CustomerApi
// @Accept json
// @Param id path int true "id主鍵"
// @Success 200 {object} common.common{}
// @Router /customer/{id} [get]
func (c *CustomerApi) GetByIdApi(ctx *gin.Context) {
	c.BaseApi.GetByIdApi(ctx)
}

// FindWithPagerApi
// @Summary Customer分頁列表
// @Description Customer分頁列表
// @Tags CustomerApi
// @Accept json
// @Param data query common.PageInfo false "頁碼,單頁大小"
// @Param T query models.Customer false "Customer條件"
// @Success 200 {object} common.common{}
// @Router /customers [get]
func (c *CustomerApi) FindWithPagerApi(ctx *gin.Context) {
	var param = common.NewSearchDto[models.Customer]()
	err := ctx.ShouldBind(param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	//排序
	//m := make(map[string]string)
	//m["created"] = "asc"
	//param.OrderRule.OrderBy = m
	withPager, i, err := c.iService.FindWithPager(*param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.PageResult(common.ErrCodeSuccess, withPager, i, ctx)
}
