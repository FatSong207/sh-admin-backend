package Repostories

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/models"
)

type MenuRepostory struct {
	Core.BaseRepostory[models.Menu, models.MenuOutDto]
}

// NewMenuRepostory CTOR
func NewMenuRepostory() IRepostories.IMenuRepostory {
	ins := &MenuRepostory{}
	return ins
}
