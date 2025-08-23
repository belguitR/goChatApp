package models

type User struct {
	Name     string    `json:"name"`
	Messages []Message `json:"messages"`
}

func newUs