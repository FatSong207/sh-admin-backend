package api

import (
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/services"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"SH-admin/taskjob"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type TaskJobApi struct {
	*BaseApi[models.TaskJob, models.TaskJob]
	iService     interface_services.ITaskJobService
	_cronService *taskjob.CronService
}

func NewTaskJobApi() *TaskJobApi {
	ins := &TaskJobApi{
		BaseApi:      NewBaseApi[models.TaskJob, models.TaskJob](),
		iService:     services.NewTaskJobService(),
		_cronService: taskjob.NewCronService(),
	}
	return ins
}

func (t *TaskJobApi) FindWithPagerApi(ctx *gin.Context) {
	var param = common.NewSearchDto[models.TaskJob]()
	//ShouldBindQuery：把query string binding到struct，struct裡面的tag要用form:"xxx"
	//ShouldBindJSON：把POST Body binding到struct，struct裡面的tag要用json:"xxx"
	err := ctx.ShouldBind(param) //ShouldBind必須在目標結構體給定form標籤
	//err := ctx.ShouldBindQuery(param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	withPager, i, err := t.iService.FindWithPager(*param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.PageResult(common.ErrCodeSuccess, withPager, i, ctx)
}

func (t *TaskJobApi) ChangStatus(ctx *gin.Context) {
	t2 := new(struct {
		Status int   `json:"status" form:"status"`
		Id     int64 `gorm:"primaryKey"`
	})
	err := ctx.ShouldBind(t2)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	m := map[string]any{"Status": t2.Status}
	//更新DB
	update, err := t.iService.Update(&models.TaskJob{Id: t2.Id}, m, false)
	if err != nil {
		common.Result(common.ErrCodeUpdateFailed, nil, ctx)
		return
	}
	if update == 0 {
		common.Result(common.ErrCodeUpdateFailed, nil, ctx)
		return
	}
	//開啟或關閉
	err = t._cronService.UpdateStatus(t2.Id, t2.Status)
	if err != nil {
		common.ResultWithMessage(common.ErrCodeUpdateFailed, "更新cron失敗"+err.Error(), nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, nil, ctx)
}

func (t *TaskJobApi) UpdateApi(ctx *gin.Context) {
	tj := new(models.TaskJob)
	err := ctx.ShouldBind(tj)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	temp := struct {
		Id   int64  `gorm:"primaryKey"`
		Cron string `json:"cron" form:"cron"`
	}{
		tj.Id,
		tj.Cron,
	}
	m := structs.Map(&temp)
	update, err := t.iService.Update(&models.TaskJob{Id: tj.Id}, m, false)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	if update == 0 {
		common.Result(common.ErrCodeUpdateFailed, nil, ctx)
		return
	}

	//更新cron表達式
	err = t._cronService.UpdateCron(temp.Id, temp.Cron)
	if err != nil {
		common.ResultWithMessage(common.ErrCodeUpdateFailed, "更新cron失敗"+err.Error(), nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, nil, ctx)

}
