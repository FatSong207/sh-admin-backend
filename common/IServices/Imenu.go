package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IMenuService interface {
	Core.IService[models.Menu, models.MenuOutDto2]
	GetMenuTree(roleId int64) (menus []models.MenuOutDto, err error)
	GetAllMenuTree(isCas bool) (menus []models.MenuOutDto, err error)
}
