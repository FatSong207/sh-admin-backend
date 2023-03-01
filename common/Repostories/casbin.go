package Repostories

import (
	"SH-admin/common/IRepostories"
)

type CasbinRepostory struct {
}

// NewCasbinRepostory CTOR
func NewCasbinRepostory() IRepostories.ICasbinRepostory {
	ins := &CasbinRepostory{}
	return ins
}
