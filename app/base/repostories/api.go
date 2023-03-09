package repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/models"
)

type ApiRepostory struct {
	core.BaseRepostory[models.Api, models.ApiOutDto]
}

// NewApiRepostory CTOR
func NewApiRepostory() interface_repostories.IApiRepostory {
	ins := &ApiRepostory{}
	return ins
}
