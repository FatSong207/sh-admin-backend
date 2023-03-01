package IRepostories

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IRoleRepostory interface {
	Core.IRepostory[models.Role, models.RoleOutDto]
}
