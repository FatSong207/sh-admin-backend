package IRepostories

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IUserRepostory interface {
	Core.IRepostory[models.User, models.UserOutDto]
}
