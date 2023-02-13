package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/models"
)

type LogApi struct {
	*BaseApi[models.Log, models.LogOutDto]
	iService IServices.ILogService
}

func NewLogApi() *LogApi {
	ins := &LogApi{
		NewBaseApi[models.Log, models.LogOutDto](),
		Services.NewLogService(),
	}
	return ins
}
