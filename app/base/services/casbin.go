package services

import (
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/repostories"
	"SH-admin/app/models"
	"SH-admin/global"
	"errors"
)

type CasbinService struct {
	_casbinRepo interface_repostories.ICasbinRepostory
	_userRepo   interface_repostories.IUserRepostory
}

// NewCasbinService 供Casbin層調用
func NewCasbinService() interface_services.ICasbinService {
	ins := &CasbinService{
		_casbinRepo: repostories.NewCasbinRepostory(),
		_userRepo:   repostories.NewUserRepostory(),
	}
	return ins
}

func (c *CasbinService) GetAccessApiByRoleId(roleId string) [][]string {
	return global.CachedEnforcer.GetFilteredPolicy(0, roleId)
}

func (c *CasbinService) UpdateCasbin(updateParam *models.UpdateCasbinParam) error {
	policy, err := global.CachedEnforcer.RemoveFilteredPolicy(0, updateParam.RoleId)
	if err != nil {
		return err
	}
	if policy {
		r := make([][]string, 0)
		for _, info := range updateParam.CasbinInfo {
			r = append(r, []string{updateParam.RoleId, info.Path, info.Method})
		}
		success, err := global.CachedEnforcer.AddPolicies(r)
		if err != nil {
			return err
		}
		if !success {
			return errors.New("存在相同api，添加失敗")
		}
		err = global.CachedEnforcer.InvalidateCache()
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("刪除原資料失敗")
}

func (c *CasbinService) UpdateUserRole(oldRules []string, newRoles []string) error {
	_, err := global.CachedEnforcer.UpdateGroupingPolicy(oldRules, newRoles)
	if err != nil {
		return err
	}
	err = global.CachedEnforcer.SavePolicy()
	if err != nil {
		return err
	}
	err = global.CachedEnforcer.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}
