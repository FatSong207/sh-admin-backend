package router

import (
	"SH-admin/api"
	"SH-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoleAuthorizeRouter(g *gin.RouterGroup) {
	rag := g.Group("/roleauthorize").Use(middleware.DbLogHandler())
	{
		rag.GET("/:roleid", api.NewRoleAuthorizeApi().GetAuthorizeIds)
		rag.PUT("", api.NewRoleAuthorizeApi().UpdateApi)
		//rag.POST("", api.NewRoleAuthorizeApi().InsertApi)
		//rag.PUT("", api.NewRoleAuthorizeApi().UpdateApi)
		//rag.DELETE("/:ids", api.NewRoleAuthorizeApi().DeleteApi)
	}
}
