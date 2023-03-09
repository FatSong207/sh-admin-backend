package api

import (
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/services"

	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"github.com/gin-gonic/gin"
)

type LogApi struct {
	*BaseApi[models.Log, models.LogOutDto]
	iService interface_services.ILogService
}

func NewLogApi() *LogApi {
	ins := &LogApi{
		NewBaseApi[models.Log, models.LogOutDto](),
		services.NewLogService(),
	}
	return ins
}

func (l *LogApi) FindWithPagerApi(ctx *gin.Context) {
	var param = common.NewSearchDto[models.Log]()
	//ShouldBindQuery：把query string binding到struct，struct裡面的tag要用form:"xxx"
	//ShouldBindJSON：把POST Body binding到struct，struct裡面的tag要用json:"xxx"
	err := ctx.ShouldBind(param) //ShouldBind必須在目標結構體給定form標籤
	//err := ctx.ShouldBindQuery(param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	//param.Entity.Type = "normalOp"
	withPager, i, err := l.iService.FindWithPager(*param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.PageResult(common.ErrCodeSuccess, withPager, i, ctx)
}

func (l *LogApi) FindLoginlogWithPagerApi(ctx *gin.Context) {
	var param = common.NewSearchDto[models.Log]()
	//ShouldBindQuery：把query string binding到struct，struct裡面的tag要用form:"xxx"
	//ShouldBindJSON：把POST Body binding到struct，struct裡面的tag要用json:"xxx"
	err := ctx.ShouldBind(param) //ShouldBind必須在目標結構體給定form標籤
	//err := ctx.ShouldBindQuery(param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	//param.Entity.Type = "normalOp"
	withPager, i, err := l.iService.FindLoginlogWithPager(*param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.PageResult(common.ErrCodeSuccess, withPager, i, ctx)
}
