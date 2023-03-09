package repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/models"
)

type MenuRepostory struct {
	core.BaseRepostory[models.Menu, models.MenuOutDto2]
}

// NewMenuRepostory CTOR
func NewMenuRepostory() interface_repostories.IMenuRepostory {
	ins := &MenuRepostory{}
	return ins
}
