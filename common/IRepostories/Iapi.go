package IRepostories

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IApiRepostory interface {
	Core.IRepostory[models.Api, models.ApiOutDto]
}
