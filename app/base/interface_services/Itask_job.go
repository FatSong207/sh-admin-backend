package interface_services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type ITaskJobService interface {
	core.IBaseService[models.TaskJob, models.TaskJob]
}
