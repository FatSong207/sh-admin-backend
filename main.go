package main

import (
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

	//time.Sleep(5 * time.Second)
	//cron1 := global.Cron["PrintMsg"]
	//cron1.Stop()
	//
	//global.Cron["PrintMsg2"] = cron.New()
	//global.Cron["PrintMsg2"].AddJob("@every 2s", taskjob.PrintMsg2{})
	//global.Cron["PrintMsg2"].Start()

	//time.Sleep(2 * time.Second)
	//global.Cron.AddJob("@every 2s", taskjob.PrintMsg2{})
	//time.Sleep(10 * time.Second)
	//global.Cron.Stop()
	//time.Sleep(3 * time.Second)
	//go initializer.InitCron()

	//claims := utils.CreateClaims(999)
	//token, err := utils.CreateToken(claims)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(token)
	//parseToken, _ := utils.ParseToken(token)
	//fmt.Println(parseToken)

	//test
	//global.Rdb.SetNX(context.Background(), "test1", "value1", 60*time.Second)

	initializer.InitRouter()
}
