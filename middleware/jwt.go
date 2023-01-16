package middleware

import (
	response "SH-admin/models/common"
	"SH-admin/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			response.Result(response.ErrCodeNoLogin, nil, ctx)
			ctx.Abort()
			return
		}
		_, err := utils.ParseToken(token)
		if err != nil {
			if err.Error() == "token is expired" {
				response.Result(response.ErrCodeTokenExpire, nil, ctx)
				ctx.Abort()
				return
			}
			response.Result(response.ErrCodeTokenError, nil, ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
