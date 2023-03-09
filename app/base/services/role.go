package services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/repostories"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"SH-admin/global"
)

type RoleService struct {
	core.IBaseService[models.Role, models.RoleOutDto]
	_roleRepo interface_repostories.IRoleRepostory
}

// NewRoleService 供Role層調用
func NewRoleService() interface_services.IRoleService {
	ins := &RoleService{
		IBaseService: core.NewBaseService[models.Role, models.RoleOutDto](),
		_roleRepo:    repostories.NewRoleRepostory(),
	}
	return ins
}

func (r *RoleService) FindWithPager(searchDto common.SearchDto[models.Role]) (*[]*models.RoleOutDto, int64, error) {
	var query = searchDto.Entity
	var dest = make([]*models.RoleOutDto, 0)
	var bind = make([]*models.Role, 0)
	var o = ""
	for k, i := range searchDto.OrderRule.OrderBy {
		o += k + " " + i
	}
	//t := new(T)
	db := global.Db.Model(query)
	if query.Name != "" {
		db = db.Where("name like ?", "%"+query.Name+"%")
	}

	t, err := r._roleRepo.FindWithPager(searchDto.PageInfo, db, o, &dest, &bind)
	if err != nil {
		return nil, 0, err
	}
	return &dest, t, nil
}
