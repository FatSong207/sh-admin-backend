package websocket

type SendMsg struct {
	Type    int    `json:"type"`
	From    string `json:"from"`
	Content string `json:"content"`
	Time    int64  `json:"time"`
}
