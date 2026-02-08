package dto

type SendMessageRequest struct {
	Sender     *UserDto `json:"sender"`
	ReceiverID int `json:"receiverId"`
	Content    string   `json:"content"`
}

type SendMessageResponse struct {
	SenderId   int `json:"senderId"`
	ReceiverID int `json:"receiverId"`
	Content    string `json:"content"`
}

type GetMessageResponse struct {
	SenderID  int `json:"senderId"`
	Content string `json:"content"`
}
