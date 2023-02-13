package Repostories

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/models"
)

type LogRepostory struct {
	Core.BaseRepostory[models.Log, models.LogOutDto]
}

// NewLogRepostory CTOR
func NewLogRepostory() IRepostories.ILogRepostory {
	ins := &LogRepostory{}
	return ins
}
