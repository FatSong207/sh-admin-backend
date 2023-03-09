package router

import (
	"SH-admin/app/api"
	"SH-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoleAuthorizeRouter(g *gin.RouterGroup) {
	api := api.NewRoleAuthorizeApi()
	rag := g.Group("/roleauthorize").Use(middleware.DbLogHandler())
	{
		rag.GET("/:roleid", api.GetAuthorizeIds)
		rag.PUT("", api.UpdateApi)
		//rag.POST("", api).InsertApi)
		//rag.PUT("", api).UpdateApi)
		//rag.DELETE("/:ids", api).DeleteApi)
	}
}
