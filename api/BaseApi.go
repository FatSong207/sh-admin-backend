package api

import (
	"SH-admin/common/Core"
	"SH-admin/models"
	response "SH-admin/models/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BaseApi[T models.Entity, TODto any] struct {
	baseSvc Core.IService[T, TODto]
}

func NewBaseApi[T models.Entity, TODto any]() *BaseApi[T, TODto] {
	ins := &BaseApi[T, TODto]{
		baseSvc: Core.NewBaseService[T, TODto](),
	}
	return ins
}

// GetById 根據Id獲取實體對應的OutDto
func (b *BaseApi[T, TODto]) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	i, _ := strconv.ParseInt(id, 10, 64)
	//getById, err := b.baseSvc.GetById(i)
	getById, err := b.baseSvc.GetOutDtoById(i)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, getById, ctx)
}

// Insert 新增實體
func (b *BaseApi[T, TODto]) Insert(ctx *gin.Context) {
	var t = new(T)
	err := ctx.ShouldBind(t)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	err, i := b.baseSvc.Insert(t, true)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, i, ctx)
}

func (b *BaseApi[T, TODto]) FindWithPager(ctx *gin.Context) {
	var param = models.NewSearchDto[T]()
	//ShouldBindQuery：把query string binding到struct，struct裡面的tag要用form:"xxx"
	//ShouldBindJSON：把POST Body binding到struct，struct裡面的tag要用json:"xxx"
	err := ctx.ShouldBind(param) //ShouldBind必須在目標結構體給定form標籤
	//err := ctx.ShouldBindQuery(param)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	withPager, i, err := b.baseSvc.FindWithPager(*param)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.PageResult(response.ErrCodeSuccess, withPager, i, ctx)
}