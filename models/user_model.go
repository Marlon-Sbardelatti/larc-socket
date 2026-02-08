package models

type User struct {
	ID       int `json:"id"`
	Password string  `json:"password"`
}

func NewUser(id int , password string) *User {
	return &User{
		ID:       id,
		Password: password,
	}
}
