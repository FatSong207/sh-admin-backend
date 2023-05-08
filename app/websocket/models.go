package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

type ClientManager struct {
	Clients map[string]*Client
	//Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

type SendMsg struct {
	Type    int    `json:"type"`
	From    string `json:"from"`
	Content string `json:"content"`
	Time    int    `json:"time"`
}

type BroadCastMsg struct {
	msgChan chan *SendMsg
}

//var msgChan = make(chan *SendMsg)

type ReplyMsg struct {
	From    string `json:"from"`
	Content string `json:"content"`
}

var Manager = ClientManager{
	Clients: make(map[string]*Client),
	//Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

var BroadCast = BroadCastMsg{
	msgChan: make(chan *SendMsg),
}

func WsHandler(c *gin.Context) {
	uid := c.Param("uid")
	//toUid := c.Query("toUid")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CORS
			return true
		}}).Upgrade(c.Writer, c.Request, nil) // 升级成ws
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// 创建一个用户实例
	client := &Client{
		ID:     uid,
		SendID: uid,
		Socket: conn,
		Send:   make(chan []byte),
	}
	// 用户注册到用户管理上
	Manager.Register <- client
	sm := &SendMsg{
		Type:    0,
		From:    "sys",
		Content: fmt.Sprintf("%s已連線", uid),
		Time:    int(time.Now().Unix()),
	}
	BroadCast.msgChan <- sm

	// 線上用戶
	BroadCast.msgChan <- &SendMsg{
		Type:    1,
		From:    "sys",
		Content: strings.Join(GetConnectedUserNames(), ","),
		Time:    int(time.Now().Unix()),
	}

	go client.Read()
	//go client.Write()
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()
	for true {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(&sendMsg) // 读取json格式，如果不是json格式，会报错
		if err != nil {
			log.Println("格式錯誤！", err)
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		if sendMsg.Type == 0 {
			sendMsg.Time = int(time.Now().Unix())
			BroadCast.msgChan <- sendMsg
			fmt.Printf("%v：%s。\n", c.ID, sendMsg.Content)
		} else if sendMsg.Type == 2 {
			disConMsg := new(SendMsg)
			disConMsg.From = "sys"
			disConMsg.Content = fmt.Sprintf("%v已離開聊天", c.ID)
			disConMsg.Time = int(time.Now().Unix())
			BroadCast.msgChan <- disConMsg
			break
		}
	}

}

func GetConnectedUserNames() []string {
	conUsers := make([]string, 0)
	for _, c2 := range Manager.Clients {
		conUsers = append(conUsers, c2.ID)
	}
	//BroadCast.msgChan <- &SendMsg{
	//	Type:    1,
	//	From:    "sys",
	//	Content: strings.Join(conUsers, ","),
	//}
	fmt.Println(conUsers)
	return conUsers
}
