package server

import (
	"github.com/belguitR/goChatApp/models"
)

type Client struct {
	Hub  *Hub
	User *models.User
}
