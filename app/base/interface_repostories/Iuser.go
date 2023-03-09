package interface_repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type IUserRepostory interface {
	core.IBaseRepostory[models.User, models.UserOutDto]
}
