package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/models"
)

type LogService struct {
	Core.IService[models.Log, models.LogOutDto]
	logRepo IRepostories.ILogRepostory
}

// NewLogService 供api層調用
func NewLogService() IServices.ILogService {
	ins := &LogService{
		IService: Core.NewBaseService[models.Log, models.LogOutDto](),
		logRepo:  Repostories.NewLogRepostory(),
	}
	return ins
}
