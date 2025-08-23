package models

type User struct {
	Name     string    `json:"name"`
	Messages []Message `json:"messages"`
}

func newUser(name string) *user {
	return &User{
		Name: name,
	}
}

func (u *User) AddMessage(m Message) {
	u.Message = append(u.messages, m)
}
