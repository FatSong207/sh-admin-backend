package initializer

import (
	"SH-admin/global"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func InitCasbin() {
	a, _ := gormadapter.NewAdapterByDB(global.DB())
	ce, _ := casbin.NewCachedEnforcer("./model.conf", a)
	ce.SetExpireTime(60 * 60)
	_ = ce.LoadPolicy()
	global.CachedEnforcer = ce
}
