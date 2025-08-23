package server

import (
	"github.com/belguitR/goChatApp/models"
	"github.com/gorilla/websocket"

	"net/http"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	User *models.User
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := &Client{
		Hub:  hub,
		Conn: conn,
		User: &models.User{Name: "Guest"},
	}
	hub.Register(client)
	defer hub.Unregister(client)
	go client.readLoop()

}

func (c *Client) readLoop() {
	defer c.Conn.Close()

	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		c.Hub.Messages = append(c.Hub.Messages, &models.Message{
			Content: string(data),
			User:    c.User,
		})

		c.Hub.Broadcast(c, data)
	}
}
