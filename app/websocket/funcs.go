package websocket

import (
	"fmt"
	"strings"
	"time"
)

// GetConnectedUserNames 獲取線上用戶
func GetConnectedUserNames(hub *hub) []string {
	conUsers := make([]string, 0)
	for _, c2 := range hub.Clients {
		conUsers = append(conUsers, c2.ID)
	}
	fmt.Println(conUsers)
	return conUsers
}

// BroadcastUsersToAll 將線上用戶廣播給大家
func BroadcastUsersToAll(c *Client) {
	c.Hub.Broadcast <- &SendMsg{
		Type:    1,
		From:    "sys",
		Content: strings.Join(GetConnectedUserNames(c.Hub), ","),
		Time:    time.Now().UnixMilli(),
	}
}
