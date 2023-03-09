package repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/models"
)

type RoleAuthorizeRepostory struct {
	core.BaseRepostory[models.RoleAuthorize, models.RoleAuthorize]
}

// NewRoleAuthorizeRepostory CTOR
func NewRoleAuthorizeRepostory() interface_repostories.IRoleAuthorizeRepostory {
	ins := &RoleAuthorizeRepostory{}
	return ins
}
