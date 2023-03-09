package interface_repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type IApiRepostory interface {
	core.IBaseRepostory[models.Api, models.ApiOutDto]
}
