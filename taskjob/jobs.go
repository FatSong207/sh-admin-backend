package taskjob

import "github.com/robfig/cron/v3"

var Jobs = []cron.Job{
	DeleteLog{},
	PrintMsg{},
	PrintMsg2{},
}
