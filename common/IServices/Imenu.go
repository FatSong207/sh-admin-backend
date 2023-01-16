package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IMenuService interface {
	Core.IService[models.Menu, models.MenuOutDto]
	GetMenuTree(roleId int64) (menus []models.MenuOutDto, err error)
	GetAllMenuTree() (menus []models.MenuOutDto, err error)
}
