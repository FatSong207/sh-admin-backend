package taskjob

import (
	"SH-admin/app/base/services"
	"fmt"
	"time"
)

type DeleteLog struct {
}

var _logService = services.NewLogService()

func (d DeleteLog) Run() {
	start := time.Now()

	_, err := _logService.DeleteAll()
	if err != nil {
		panic(err)
	}
	cost := time.Since(start)
	fmt.Println(cost)
}

type PrintMsg struct {
}

func (p PrintMsg) Run() {
	fmt.Println("msg")
}

type PrintMsg2 struct {
}

func (p PrintMsg2) Run() {
	fmt.Println("msg22")
}
