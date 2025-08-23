package models

type User struct {
	Name string `json:"name"`
}

func newUser(name string) *User {
	return &User{
		Name: name,
	}
}
