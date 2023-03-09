package repostories

import (
	"SH-admin/app/base/interface_repostories"
)

type CasbinRepostory struct {
}

// NewCasbinRepostory CTOR
func NewCasbinRepostory() interface_repostories.ICasbinRepostory {
	ins := &CasbinRepostory{}
	return ins
}
