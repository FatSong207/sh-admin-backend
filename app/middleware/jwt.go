package middleware

import (
	"SH-admin/app/models/common"
	"SH-admin/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			common.Result(common.ErrCodeNoLogin, nil, ctx)
			ctx.Abort()
			return
		}
		_, err := utils.ParseToken(token)
		if err != nil {
			if err.Error() == "token is expired" {
				common.Result(common.ErrCodeTokenExpire, nil, ctx)
				ctx.Abort()
				return
			}
			common.Result(common.ErrCodeTokenError, nil, ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
