package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/global"
	"SH-admin/models"
	"SH-admin/models/common"
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

	t, err := r.RoleRepo.FindWithPager(searchDto.PageInfo, db, o, &dest, &bind)
	if err != nil {
		return nil, 0, err
	}
	return &dest, t, nil
}
