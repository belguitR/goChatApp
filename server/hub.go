package server

import (
	"sync"
	"log"
	"github.com/belguitR/goChatApp/models"
	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients  []*Client
	Messages []*models.Message
	mu       sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:  make([]*Client, 0),
		Messages: make([]*models.Message, 0),
	}
}

func (h *Hub) Register(c *Client) {
	h.mu.Lock()
	h.Clients = append(h.Clients, c)
	h.mu.Unlock()
}

func (h *Hub) Unregister(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	idx := h.findPosition(c)
	if idx == -1 {
		return
	}

	h.Clients = append(h.Clients[:idx], h.Clients[idx+1:]...)
}

func (h *Hub) Broadcast(sender *Client, data []byte) {
	h.mu.RLock()
	clients := make([]*Client, len(h.Clients))
	copy(clients, h.Clients)
	h.mu.RUnlock()

	for _, c := range clients {
		log.Println("Broadcasting to", c.User.Name, ":", string(data))
		err := c.Conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("failed to send to", c.User.Name, err)
			go h.Unregister(c)
		}
	}
}


func (h *Hub) findPosition(c *Client) int {
	i := 0
	for i < len(h.Clients) {
		if h.Clients[i] == c {
			return i
		}
		i = i + 1
	}
	return -1
}

func (h *Hub) OnlineCount() int {
  h.mu.RLock()
  defer h.mu.RUnlock()
  return len(h.Clients)
}

