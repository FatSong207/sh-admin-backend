package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/models"
	"SH-admin/models/common"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiApi struct {
	*BaseApi[models.Api, models.ApiOutDto]
	iService IServices.IApiService
}

func NewApiApi() *ApiApi {
	ins := &ApiApi{
		NewBaseApi[models.Api, models.ApiOutDto](),
		Services.NewApiService(),
	}
	return ins
}

func (a *ApiApi) InsertApi(ctx *gin.Context) {
	var t = new(models.Api)
	err := ctx.ShouldBind(t)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	/**CHECK**/
	tDb, err := a.iService.GetByWhereStruct(&models.Api{Path: t.Path, Method: t.Method})
	if err != nil && err != gorm.ErrRecordNotFound {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	if tDb != nil {
		common.ResultWithMessage(common.ErrCodeInsertFailed, "已有相同路徑以及方法的資料！", nil, ctx)
		return
	}

	err, i := a.iService.Insert(t, false)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, i, ctx)
}

func (a *ApiApi) UpdateApi(ctx *gin.Context) {
	api := new(models.ApiInDto)
	err := ctx.ShouldBind(api)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	/**CHECK**/
	tDb, err := a.iService.GetByWhereStruct(&models.Api{Path: api.Path, Method: api.Method})
	if err != nil && err != gorm.ErrRecordNotFound {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	if tDb != nil {
		common.ResultWithMessage(common.ErrCodeInsertFailed, "已有相同路徑以及方法的資料！", nil, ctx)
		return
	}

	am := structs.Map(api)
	update, err := a.iService.Update(&models.Api{Id: api.Id}, am, false)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	if update == 0 {
		common.Result(common.ErrCodeUpdateFailed, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, update, ctx)
}
