package middleware

import (
	"SH-admin/global"
	"SH-admin/models/common"
	"SH-admin/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AuthorizeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := utils.GetClaims(ctx)
		if err != nil {
			ctx.Abort()
			return
		}
		sub := claims.Uid
		obj := ctx.FullPath()
		act := ctx.Request.Method
		if claims.RoleId == 12 {
			ctx.Next()
		} else {
			success, _ := global.CachedEnforcer.Enforce(strconv.FormatInt(sub, 10), obj, act)
			if success {
				ctx.Next()
			} else {
				common.Result(common.ErrCode403, nil, ctx)
				ctx.Abort()
			}
		}

	}
}
