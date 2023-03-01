package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IRoleAuthorizeService interface {
	Core.IService[models.RoleAuthorize, models.RoleAuthorize]
	UpdateBatchByRoleId(roleId int64, authorizeIds []int64) (AffectedRows int64, err error)
}
