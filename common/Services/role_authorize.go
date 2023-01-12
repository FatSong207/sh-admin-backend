package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/models"
)

type RoleAuthorizeService struct {
	Core.IService[models.RoleAuthorize, models.RoleAuthorize]
	roleAuthRepo IRepostories.IRoleAuthorizeRepostory
}

// NewRoleAuthorizeService 供api層調用
func NewRoleAuthorizeService() IServices.IRoleAuthorizeService {
	ins := &RoleAuthorizeService{
		IService:     Core.NewBaseService[models.RoleAuthorize, models.RoleAuthorize](),
		roleAuthRepo: Repostories.NewRoleAuthorizeRepostory(),
	}
	return ins
}
