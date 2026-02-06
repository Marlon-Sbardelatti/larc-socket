package main

import "fmt"

type Message struct {
	method  string
	msgType string
	content string
}

func NewMessage(method string, msgType string, content string) *Message {
	return &Message{
		method:  method,
		msgType: msgType,
		content: content,
	}
}

func (m *Message) buildHeaders() string {
	return fmt.Sprintf("%s %s", m.method, m.content)
}

func BuildPayload(message Message, user User) string {
	payload := message.buildHeaders() + " " + user.String() + message.content + ":" + message.content

	return payload
}
