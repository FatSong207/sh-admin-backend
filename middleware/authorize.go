package middleware

import (
	response "SH-admin/models/common"
	"github.com/gin-gonic/gin"
)

func AuthorizeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response.Result(response.ErrCode403, nil, ctx)
		ctx.Abort()
		return
	}
}
