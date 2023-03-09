package interface_repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type IRoleAuthorizeRepostory interface {
	core.IBaseRepostory[models.RoleAuthorize, models.RoleAuthorize]
}
