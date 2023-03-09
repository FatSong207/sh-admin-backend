package services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/repostories"
	"SH-admin/app/models"
)

type RoleAuthorizeService struct {
	core.IBaseService[models.RoleAuthorize, models.RoleAuthorize]
	_roleAuthRepo interface_repostories.IRoleAuthorizeRepostory
}

// NewRoleAuthorizeService 供api層調用
func NewRoleAuthorizeService() interface_services.IRoleAuthorizeService {
	ins := &RoleAuthorizeService{
		IBaseService:  core.NewBaseService[models.RoleAuthorize, models.RoleAuthorize](),
		_roleAuthRepo: repostories.NewRoleAuthorizeRepostory(),
	}
	return ins
}

func (r *RoleAuthorizeService) UpdateBatchByRoleId(roleId int64, authorizeIds []int64) (AffectedRows int64, err error) {
	_, err = r._roleAuthRepo.DeleteByKeys([]int{int(roleId)})
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
	err, i := r._roleAuthRepo.InsertBatch(&l, false)
	if err != nil {
		return 0, err
	}
	return i, err
}
