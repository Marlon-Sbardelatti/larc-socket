package dto

import (
	"main.go/models"
)

type SendMessageRequest struct {
	Sender     models.User `json:"sender"`
	ReceiverID string      `json:"receiverId"`
	Content    string      `json:"content"`
}

type GetMessageRequest struct {
	User models.User `json:"user"`
}

type GetMessageResponse struct {
	UserID string
	Content string
}
