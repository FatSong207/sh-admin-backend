package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IRoleService interface {
	Core.IService[models.Role, models.RoleOutDto]
}
