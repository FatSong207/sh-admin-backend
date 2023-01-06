package Repostories

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/models"
)

type UserRepostory struct {
	Core.BaseRepostory[models.User, models.UserOutDto]
}

// NewUserRepostory CTOR
func NewUserRepostory() IRepostories.IUserRepostory {
	ins := &UserRepostory{}
	return ins
}
