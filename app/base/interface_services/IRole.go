package interface_services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
)

type IRoleService interface {
	core.IBaseService[models.Role, models.RoleOutDto]
	FindWithPager(searchDto common.SearchDto[models.Role]) (*[]*models.RoleOutDto, int64, error)
}
