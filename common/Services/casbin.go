package Services

import (
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/global"
	"SH-admin/models"
	"errors"
)

type CasbinService struct {
	CasbinRepo IRepostories.ICasbinRepostory
	userRepo   IRepostories.IUserRepostory
}

// NewCasbinService 供Casbin層調用
func NewCasbinService() IServices.ICasbinService {
	ins := &CasbinService{
		CasbinRepo: Repostories.NewCasbinRepostory(),
		userRepo:   Repostories.NewUserRepostory(),
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
