package global

import (
	"SH-admin/config"
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config         config.Config
	Rdb            *redis.Client
	db             *gorm.DB
	Log            *zap.SugaredLogger
	CachedEnforcer *casbin.CachedEnforcer
	Cron           map[string]*cron.Cron
)

// DB 獲取db
func DB() *gorm.DB {
	return db
}

// SetDB 初始化時設置db
func SetDB(gdb *gorm.DB) {
	db = gdb
}
