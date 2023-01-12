package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IMenuService interface {
	Core.IService[models.Menu, models.MenuOutDto]
	GetMenuTreeMap(roleId int64) (treeMap map[string][]models.MenuOutDto, err error)
	GetMenuTree(roleId int64) (menus []models.MenuOutDto, err error)
}
