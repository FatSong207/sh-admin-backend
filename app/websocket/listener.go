package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
)

func Listen() {
	for true {
		select {
		//監聽是否有人進入聊天室
		case conn := <-Manager.Register:
			fmt.Println(conn.ID + "已連線")
			Manager.Clients[conn.ID] = conn
		//監聽訊息的發送
		case bmsg := <-BroadCast.msgChan:
			msg, _ := json.Marshal(bmsg)
			for _, client := range Manager.Clients {
				_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		//監聽是否有人離開聊天室
		case conn := <-Manager.Unregister:
			if _, ok := Manager.Clients[conn.ID]; ok {
				fmt.Println(conn.ID, "已離開聊天")
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
				// 線上用戶
				sm1 := new(SendMsg)
				sm1.Type = 1
				sm1.From = "sys"
				sm1.Content = strings.Join(GetConnectedUserNames(), ",")
				go func() {
					BroadCast.msgChan <- sm1
				}()
				fmt.Println(sm1)
			}
		}
	}
}
