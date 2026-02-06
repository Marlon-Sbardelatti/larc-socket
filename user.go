package main

import "fmt"

type User struct {
	id string
	password string
}

func NewUser(id string, password string) *User {
	return &User{
		id: id,
		password: password,
	}
}

func (m *User) String() string {
	return fmt.Sprintf("%s:%s", m.id, m.password)
}
