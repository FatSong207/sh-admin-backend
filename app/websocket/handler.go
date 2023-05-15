package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func WsHandler(c *gin.Context) {
	uid := c.Param("uid")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CORS
			return true
		}}).Upgrade(c.Writer, c.Request, nil) // 升级成ws
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//新建客戶端
	client := &Client{
		ID:     uid,
		Socket: conn,
		Hub:    Hub,
		Send:   make(chan *SendMsg, 1024),
	}
	Hub.Register <- client //註冊進hub中

	go client.Read()
	go client.Write()
}
