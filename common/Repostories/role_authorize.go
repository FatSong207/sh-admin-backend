package Repostories

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/models"
)

type RoleAuthorizeRepostory struct {
	Core.BaseRepostory[models.RoleAuthorize, models.RoleAuthorize]
}

// NewRoleAuthorizeRepostory CTOR
func NewRoleAuthorizeRepostory() IRepostories.IRoleAuthorizeRepostory {
	ins := &RoleAuthorizeRepostory{}
	return ins
}
