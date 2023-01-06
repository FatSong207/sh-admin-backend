package main

import (
	_ "SH-admin/docs"
	"SH-admin/global"
	"SH-admin/initializer"
	"context"
	"time"
)

// @title sh-Admin Api
// @version 1.0
// @description sh-Admin 後端接口

// @host localhost:5001
// @BasePath /api

func main() {
	initializer.InitViper()

	initializer.InitLogger()
	defer global.Log.Sync()

	initializer.InitRedis()
	initializer.InitGorm()

	//claims := utils.CreateClaims(999)
	//token, err := utils.CreateToken(claims)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(token)
	//parseToken, _ := utils.ParseToken(token)
	//fmt.Println(parseToken)

	//test
	global.Rdb.SetNX(context.Background(), "test1", "value1", 60*time.Second)

	initializer.InitRouter()
}
