package interface_repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type ITaskJobRepostory interface {
	core.IBaseRepostory[models.TaskJob, models.TaskJob]
}
