package interface_repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type IRoleRepostory interface {
	core.IBaseRepostory[models.Role, models.RoleOutDto]
}
