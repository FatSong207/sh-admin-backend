package initializer

import (
	"SH-admin/global"
	"fmt"
	"github.com/go-redis/redis/v9"
)

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
	})
	global.Rdb = rdb
}
