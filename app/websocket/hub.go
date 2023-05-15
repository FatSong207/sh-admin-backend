package websocket

import (
	"fmt"
	"time"
)

type hub struct {
	Clients    map[string]*Client
	Broadcast  chan *SendMsg
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

var Hub = newHub()

func newHub() *hub {
	return &hub{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan *SendMsg, 1024),
		Reply:      make(chan *Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *hub) Run() {
	for true {
		select {
		case c := <-h.Register:
			fmt.Printf("\n%v已連線", c.ID)
			h.Clients[c.ID] = c

			//已連線資訊廣播給大家
			sm := &SendMsg{
				Type:    0,
				From:    "sys",
				Content: fmt.Sprintf("%s已連線", c.ID),
				Time:    time.Now().UnixMilli(),
			}
			Hub.Broadcast <- sm

			// 線上用戶廣播給大家
			BroadcastUsersToAll(c)
		case c := <-h.Unregister:
			fmt.Printf("\n%v已離開", c.ID)
			close(c.Send)
			delete(h.Clients, c.ID)
			// 線上用戶廣播給大家
			BroadcastUsersToAll(c)
		case msg := <-h.Broadcast:
			fmt.Printf("%v：%s。\n", msg.From, msg.Content)
			for _, client := range h.Clients {
				client.Send <- msg
			}
		}
	}
}
