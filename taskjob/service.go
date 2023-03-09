package taskjob

import (
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/services"
	"SH-admin/global"
	"github.com/robfig/cron/v3"
	"reflect"
	"strconv"
)

type CronService struct {
	_taskjobService interface_services.ITaskJobService
}

func NewCronService() *CronService {
	return &CronService{
		_taskjobService: services.NewTaskJobService(),
	}
}

// UpdateStatus 開啟或關閉cron
func (c CronService) UpdateStatus(id int64, status int) error {
	if status == 0 { //關閉
		cronIns := global.Cron[strconv.FormatInt(id, 10)]
		cronIns.Stop()
	} else if status == 1 { //開啟
		global.Cron[strconv.FormatInt(id, 10)].Start()
	}
	return nil
}

// UpdateCron 更新cron表達式
func (c CronService) UpdateCron(id int64, cronStr string) error {
	//重啟定時任務
	if c, ok := global.Cron[strconv.FormatInt(id, 10)]; ok {
		c.Stop()
		delete(global.Cron, strconv.FormatInt(id, 10))
	}
	taskJob, err := c._taskjobService.GetById(id)
	if err != nil {
		return err
	}
	for _, job := range Jobs {
		if reflect.TypeOf(job).Name() == taskJob.TaskName {
			tj := cron.New()
			tj.AddJob(cronStr, job)
			if taskJob.Status == 1 {
				tj.Start()
				//defer tj.Stop()
			}
			global.Cron[strconv.FormatInt(id, 10)] = tj
		}
	}
	return nil
}
