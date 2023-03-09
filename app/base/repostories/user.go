package repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/models"
)

type UserRepostory struct {
	core.BaseRepostory[models.User, models.UserOutDto]
}

// NewUserRepostory CTOR
func NewUserRepostory() interface_repostories.IUserRepostory {
	ins := &UserRepostory{}
	return ins
}
