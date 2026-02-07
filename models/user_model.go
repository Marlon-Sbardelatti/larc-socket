package models

type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func NewUser(id, password string) *User {
	return &User{
		ID:       id,
		Password: password,
	}
}

