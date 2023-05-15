package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	Socket *websocket.Conn
	Hub    *hub
	Send   chan *SendMsg
}

func (c *Client) Read() {
	defer func() {
		c.Hub.Unregister <- c
		c.Socket.Close()
	}()
	for true {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(&sendMsg) // 必須為json
		if err != nil {
			log.Println("格式錯誤！", err)
			c.Hub.Unregister <- c //這段有需要嗎？ defer fun()不是做掉了？
			_ = c.Socket.Close()  //這段有需要嗎？ defer fun()不是做掉了？
			break
		}
		if sendMsg.Type == 0 {
			sendMsg.Time = time.Now().UnixMilli()
			c.Hub.Broadcast <- sendMsg
			fmt.Printf("%v：%s。\n", c.ID, sendMsg.Content)
		} else if sendMsg.Type == 2 {
			disConMsg := new(SendMsg)
			disConMsg.From = "sys"
			disConMsg.Content = fmt.Sprintf("%v已離開聊天", c.ID)
			disConMsg.Time = time.Now().UnixMilli()
			c.Hub.Broadcast <- disConMsg
			break
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Hub.Unregister <- c
		c.Socket.Close()
	}()
	for true {
		select {
		case smsg := <-c.Send:
			msg, _ := json.Marshal(smsg)
			c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
