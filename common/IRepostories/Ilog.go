package IRepostories

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type ILogRepostory interface {
	Core.IRepostory[models.Log, models.LogOutDto]
}
