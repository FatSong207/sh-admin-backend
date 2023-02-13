package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type ILogService interface {
	Core.IService[models.Log, models.LogOutDto]
}
