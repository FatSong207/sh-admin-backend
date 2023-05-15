package main

import (
	"SH-admin/app/websocket"
	_ "SH-admin/docs"
	"SH-admin/global"
	"SH-admin/initializer"
)

// @title sh-Admin Api
// @version 1.0
// @description sh-Admin 後端接口

// @host localhost:5001
// @BasePath /api
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name token

func main() {
	initializer.InitViper()

	initializer.InitLogger()
	defer global.Log.Sync()

	initializer.InitRedis()
	initializer.InitGorm()
	initializer.InitCasbin()

	sigChan := make(chan int) //信號channel
	go initializer.InitCron(sigChan)
	<-sigChan //若上方InitCron沒做完會阻塞在這邊(確保InitCron做完)

	go websocket.Hub.Run()
	initializer.InitRouter()
}
