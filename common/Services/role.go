package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/models"
)

type RoleService struct {
	Core.IService[models.Role, models.RoleOutDto]
	RoleRepo IRepostories.IRoleRepostory
}

// NewRoleService 供Role層調用
func NewRoleService() IServices.IRoleService {
	ins := &RoleService{
		IService: Core.NewBaseService[models.Role, models.RoleOutDto](),
		RoleRepo: Repostories.NewRoleRepostory(),
	}
	return ins
}
