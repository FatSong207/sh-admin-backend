package initializer

import (
	"SH-admin/api"
	"SH-admin/global"
	"SH-admin/middleware"
	"SH-admin/router"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	e := gin.Default()

	// 開啟跨域
	e.Use(middleware.Cors())

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//公共路由
	publicGroup := e.Group("/api").Use(middleware.DbLogHandler())
	{
		publicGroup.GET("health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})
		publicGroup.POST("login", api.NewUserApi().Login)
		publicGroup.GET("user/verifycode", api.NewUserApi().GetVerifyCode)
		publicGroup.POST("user/register", api.NewUserApi().Register)
	}

	//私有路由
	privateGroup := e.Group("/api")
	privateGroup.Use(middleware.LogHandler()).Use(middleware.JwtAuth())
	{
		router.InitProductRouter(privateGroup)
		router.InitCustomerRouter(privateGroup)
		router.InitUserRouter(privateGroup)
		router.InitMenuRouter(privateGroup)
		router.InitSystemRouter(privateGroup)
		router.InitLogRouter(privateGroup)
	}

	e.Run(fmt.Sprintf(":%s", global.Config.Port))
}
