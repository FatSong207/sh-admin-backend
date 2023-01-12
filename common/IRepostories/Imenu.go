package IRepostories

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IMenuRepostory interface {
	Core.IRepostory[models.Menu, models.MenuOutDto]
}
