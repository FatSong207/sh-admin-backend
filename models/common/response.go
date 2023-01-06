package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Page struct {
	Total int64 `json:"total"`
	List  any   `json:"list"`
}

func Result(code int, data any, ctx *gin.Context) {
	message := msg[code]
	ctx.JSON(http.StatusOK, Response{
		code,
		message,
		data,
	})
}

func PageResult(code int, data any, rows int64, c *gin.Context) {
	message := msg[code]
	page := &Page{Total: rows, List: data}
	c.JSON(http.StatusOK, Response{
		code,
		message,
		page,
	})
}
