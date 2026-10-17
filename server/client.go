package server

import (
	"net/http"
	"sync"

	"github.com/belguitR/goChatApp/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub     *Hub
	Conn    *websocket.Conn
	User    *models.User
	writeMu sync.Mutex
}

func (c *Client) Send(data []byte) error {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	return c.Conn.WriteMessage(websocket.TextMessage, data)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
	go client.readLoop()

}

func (c *Client) readLoop() {
	defer func() {
		c.Hub.Unregister(c)
		c.Conn.Close()
	}()

	nameMessage := true // The first message from the client is its display name.

	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		if nameMessage {
			c.User = &models.User{Name: string(data)}
			nameMessage = false
			continue
		}

		msg := c.User.Name + ": " + string(data)
		c.Hub.RecordMessage(&models.Message{
			Content: string(data),
			User:    c.User,
		})

		c.Hub.Broadcast([]byte(msg))
	}
}
