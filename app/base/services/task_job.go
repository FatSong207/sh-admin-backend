package services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/repostories"
	"SH-admin/app/models"
)

type TaskJobService struct {
	core.IBaseService[models.TaskJob, models.TaskJob]
	_taskjobRepo interface_repostories.ITaskJobRepostory
}

// NewTaskJobService 供api層調用
func NewTaskJobService() interface_services.ITaskJobService {
	ins := &TaskJobService{
		IBaseService: core.NewBaseService[models.TaskJob, models.TaskJob](),
		_taskjobRepo: repostories.NewTaskJobRepostory(),
	}
	return ins
}

//func (t *TaskJobService) UpdateStatus(id int64, status int) error {
//	//更新cron
//	if status == 0 { //關閉
//		cron := global.Cron[strconv.FormatInt(id, 10)]
//		cron.Stop()
//	} else if status == 1 { //開啟
//		id, err := t._taskjobRepo.GetById(id)
//		if err != nil {
//			return err
//		}
//		global.Cron[strconv.FormatInt(id.Id, 10)].Start()
//	}
//	return nil
//}
