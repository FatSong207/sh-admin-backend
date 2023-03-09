package interface_repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type IMenuRepostory interface {
	core.IBaseRepostory[models.Menu, models.MenuOutDto2]
}
