package initializer

import (
	"SH-admin/app/base/services"
	"SH-admin/global"
	"SH-admin/taskjob"
	"github.com/robfig/cron/v3"
	"reflect"
	"strconv"
)

var _taskjobService = services.NewTaskJobService()

func InitCron(sigChan chan int) {
	//global.Cron = cron.New()

	all, err := _taskjobService.GetAll()
	if err != nil {
		panic(err)
	}
	global.Cron = make(map[string]*cron.Cron, 10)

	for _, job := range all {
		for _, c := range taskjob.Jobs {
			if reflect.TypeOf(c).Name() == job.TaskName {
				t := cron.New()
				_, err := t.AddJob(job.Cron, c)
				if err != nil {
					panic(err)
				}
				//fmt.Println(addJob)
				if job.Status == 1 {
					t.Start()
					defer t.Stop()
				}
				global.Cron[strconv.FormatInt(job.Id, 10)] = t
			}
		}
	}
	sigChan <- 1
	//defer global.Cron.Stop()
	select {}
}
