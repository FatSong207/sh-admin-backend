package global

import (
	"SH-admin/config"
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config         config.Config
	Rdb            *redis.Client
	Db             *gorm.DB
	Log            *zap.SugaredLogger
	CachedEnforcer *casbin.CachedEnforcer
)
