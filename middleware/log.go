package middleware

import (
	"SH-admin/global"
	"github.com/gin-gonic/gin"
)

func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		global.Log.Infof("From: %s  Method:%s URL:%s", c.ClientIP(), c.Request.Method, c.Request.URL.Path+c.Request.URL.RawQuery)
		c.Next()
	}
}
