package interface_services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type IRoleAuthorizeService interface {
	core.IBaseService[models.RoleAuthorize, models.RoleAuthorize]
	UpdateBatchByRoleId(roleId int64, authorizeIds []int64) (AffectedRows int64, err error)
}
