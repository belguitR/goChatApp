package models

type Message struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	User    string `json:"user"`
	// Time will be added later
}

func NewMessage(id int, content string, user string) *Message {
	return &Message{
		Id:      id,
		Content: content,
		User:    user,
	}
}
