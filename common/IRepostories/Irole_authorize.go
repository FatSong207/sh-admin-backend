package IRepostories

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IRoleAuthorizeRepostory interface {
	Core.IRepostory[models.RoleAuthorize, models.RoleAuthorize]
}
