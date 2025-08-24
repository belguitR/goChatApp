package server

import (
	"log"
	"net/http"

	"github.com/belguitR/goChatApp/models"
	"github.com/gorilla/websocket"
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
	go client.readLoop()

}

func (c *Client) readLoop() {
	defer func() {
		c.Hub.Unregister(c)
		c.Conn.Close()
	}()

		nameMessage:= true // the first message from the clinet will always be the username

	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		if nameMessage {
			c.User = &models.User{Name : string(data)}
			nameMessage = false 
			log.Println("user  ",c.User.Name)
			continue
		}
		log.Println("server recv:", string(data))
		
		msg := c.User.Name + ": " + string(data)
		c.Hub.Messages = append(c.Hub.Messages, &models.Message{
			Content: string(data),
			User:    c.User,
		})

		c.Hub.Broadcast(c, []byte(msg))
	}
}
