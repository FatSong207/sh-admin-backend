package repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/models"
)

type roleRepostory struct {
	core.BaseRepostory[models.Role, models.RoleOutDto]
}

// NewRoleRepostory CTOR
func NewRoleRepostory() interface_repostories.IRoleRepostory {
	ins := &roleRepostory{}
	return ins
}
