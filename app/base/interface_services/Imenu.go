package interface_services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type IMenuService interface {
	core.IBaseService[models.Menu, models.MenuOutDto2]
	GetMenuTree(roleId int64) (menus []models.MenuOutDto, err error)
	GetAllMenuTree(isCas bool) (menus []models.MenuOutDto, err error)
}
