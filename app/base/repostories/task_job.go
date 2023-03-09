package repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/models"
)

type TaskJobRepostory struct {
	core.BaseRepostory[models.TaskJob, models.TaskJob]
}

// NewTaskJobRepostory CTOR
func NewTaskJobRepostory() interface_repostories.ITaskJobRepostory {
	ins := &TaskJobRepostory{}
	return ins
}
