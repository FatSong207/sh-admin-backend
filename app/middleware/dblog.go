package middleware

import (
	"SH-admin/app/base/services"
	"SH-admin/app/models"
	"SH-admin/utils"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func DbLogHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := &models.Log{}
		if ctx.Request.Method != http.MethodGet {
			b, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				ctx.Abort()
				return
			}
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(b))
			log.Body = string(b) //賦值
		} else {
			if ctx.Request.URL.RawQuery != "" {
				q := ctx.Request.URL.RawQuery
				q, _ = url.QueryUnescape(q)
				g := strings.Split(q, "&")
				m := make(map[string]string)
				for _, s := range g {
					kv := strings.Split(s, "=")
					if len(kv) == 2 {
						m[kv[0]] = kv[1]
					}
				}
				body, err := json.Marshal(&m)
				if err != nil {
					ctx.Abort()
					return
				}
				log.Body = string(body) //賦值
			}
		}

		//獲取訪問人
		claims, _ := utils.GetClaims(ctx)
		if claims != nil {
			log.UserID = int(claims.Uid) //賦值
		} else {
			log.UserID = 0 //賦值
		}

		//登入操作
		if ctx.FullPath() == "/api/login" {
			log.Type = "Login" //賦值
		} else {
			log.Type = "normalOp" //賦值
		}

		log.Path = ctx.Request.URL.Path     //賦值
		log.Agent = ctx.Request.UserAgent() //賦值
		log.Ip = ctx.ClientIP()             //賦值
		log.Method = ctx.Request.Method     //賦值

		//獲取ResponseBody
		//https://cloud.tencent.com/developer/article/1811436
		writer := responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}
		ctx.Writer = writer

		ctx.Next()

		log.Status = ctx.Writer.Status() //賦值
		//log.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String() //賦值

		body := writer.body.String()
		log.Response = body
		m := make(map[string]any)
		json.Unmarshal([]byte(body), &m)
		log.Code = int(m["code"].(float64))

		services.NewLogService().Insert(log, false)
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
