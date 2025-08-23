package server

import (
	"github.com/belguitR/goChatApp/models"
)

type Hub struct {
	Clients  []*Client
	Messages []*models.Message
}

func NewHub() *Hub {
	return &Hub{
		Clients:  make([]*Client, 0),
		Messages: make([]*models.Message, 0),
	}
}

func register(c *Client, h *Hub) {
	h.Clients = append(h.Clients, c)
}

func unregister(c *Client, h *Hub) {
	idx := h.findPosition(c)
	if idx == -1 {
		return
	}
	h.Clients = append(h.Clients[:idx], h.Clients[idx+1:]...)
}

func broadcast()

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
