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
		publicGroup.POST("login", api.NewLoginApi().Login)
		publicGroup.POST("register", api.NewLoginApi().Register)
	}
	publicGroupWithoutDbLog := e.Group("/api")
	{
		publicGroup.GET("health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})
		publicGroupWithoutDbLog.GET("verifycode", api.NewLoginApi().GetVerifyCode)
		publicGroupWithoutDbLog.GET("userinfo", api.NewLoginApi().GetUserInfoApi)
	}

	//私有路由
	privateGroup := e.Group("/api")
	privateGroup.Use(middleware.LogHandler()).Use(middleware.JwtAuth()).Use(middleware.AuthorizeHandler())
	{
		router.InitProductRouter(privateGroup)
		router.InitCustomerRouter(privateGroup)
		router.InitUserRouter(privateGroup)
		router.InitMenuRouter(privateGroup)
		router.InitSystemRouter(privateGroup)
		router.InitLogRouter(privateGroup)
		router.InitApiRouter(privateGroup)
		router.InitRoleRouter(privateGroup)
		router.InitRoleAuthorizeRouter(privateGroup)
		router.InitCasbinRouter(privateGroup)
	}

	e.Run(fmt.Sprintf(":%s", global.Config.Port))
}
