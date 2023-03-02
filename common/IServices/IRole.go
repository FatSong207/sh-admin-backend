package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
	"SH-admin/models/common"
)

type IRoleService interface {
	Core.IService[models.Role, models.RoleOutDto]
	FindWithPager(searchDto common.SearchDto[models.Role]) (*[]*models.RoleOutDto, int64, error)
}
