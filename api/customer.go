package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/models"
	response "SH-admin/models/common"
	"github.com/gin-gonic/gin"
)

type CustomerApi struct {
	*BaseApi[models.Customer, models.CustomerOutDto]
	IServices.ICustomerService
}

func NewCustomerApi() *CustomerApi {
	ins := &CustomerApi{
		NewBaseApi[models.Customer, models.CustomerOutDto](),
		Services.NewCustomerService(),
	}
	return ins
}

// GetByEmail
// @Summary 根據Email獲取Customer
// @Description 分頁列表
// @Tags CustomerApi
// @Accept json
// @Param email path string true "EMail"
// @Success 200 {object} response.Response{}
// @Router /customer/GetByEmail/{email} [get]
func (c *CustomerApi) GetByEmail(ctx *gin.Context) {
	e := ctx.Param("email")
	t, err := c.ICustomerService.GetByEmail(e)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, t, ctx)
}

// GetById @Summary 根據Id獲取Customer對應的OutDto
// @Description 根據Id獲取Customer對應的OutDto
// @Tags CustomerApi
// @Accept json
// @Param id path int true "id主鍵"
// @Success 200 {object} response.Response{}
// @Router /customer/{id} [get]
func (c *CustomerApi) GetById(ctx *gin.Context) {
	c.BaseApi.GetById(ctx)
}

// FindWithPager
// @Summary Customer分頁列表
// @Description Customer分頁列表
// @Tags CustomerApi
// @Accept json
// @Param data query response.PageInfo false "頁碼,單頁大小"
// @Param T query models.Customer false "Customer條件"
// @Success 200 {object} response.Response{}
// @Router /customers [get]
func (c *CustomerApi) FindWithPager(ctx *gin.Context) {
	var param = models.NewSearchDto[models.Customer]()
	err := ctx.ShouldBind(param)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	withPager, i, err := c.ICustomerService.FindWithPager(*param)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.PageResult(response.ErrCodeSuccess, withPager, i, ctx)
}
