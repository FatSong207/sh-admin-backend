package Repostories

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/models"
)

type RoleRepostory struct {
	Core.BaseRepostory[models.Role, models.RoleOutDto]
}

// NewRoleRepostory CTOR
func NewRoleRepostory() IRepostories.IRoleRepostory {
	ins := &RoleRepostory{}
	return ins
}
