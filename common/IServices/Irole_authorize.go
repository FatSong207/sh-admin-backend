package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IRoleAuthorizeService interface {
	Core.IService[models.RoleAuthorize, models.RoleAuthorize]
}
