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

func (r *RoleAuthorizeService) UpdateBatchByRoleId(roleId int64, authorizeIds []int64) (AffectedRows int64, err error) {
	_, err = r.roleAuthRepo.DeleteByKeys([]int{int(roleId)})
	if err != nil {
		return 0, err
	}
	l := make([]*models.RoleAuthorize, 0)
	for _, id := range authorizeIds {
		l = append(l, &models.RoleAuthorize{
			RoleId:      roleId,
			AuthorizeId: id,
		})
	}
	err, i := r.roleAuthRepo.InsertBatch(&l, false)
	if err != nil {
		return 0, err
	}
	return i, err
}
