package models

type Message struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	User    *User  `json:"user"`
	// Time will be added later
}

func NewMessage(id int, content string, user *User) *Message {
	return &Message{
		Id:      id,
		Content: content,
		User:    user,
	}
}
