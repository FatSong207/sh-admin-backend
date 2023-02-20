package Repostories

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/models"
)

type ApiRepostory struct {
	Core.BaseRepostory[models.Api, models.ApiOutDto]
}

// NewApiRepostory CTOR
func NewApiRepostory() IRepostories.IApiRepostory {
	ins := &ApiRepostory{}
	return ins
}
