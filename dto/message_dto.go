package dto

type SendMessageRequest struct {
	Sender     *UserDto `json:"sender"`
	ReceiverID string   `json:"receiverId"`
	Content    string   `json:"content"`
}

type SendMessageResponse struct {
	SenderId   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
	Content    string `json:"content"`
}

type GetMessageResponse struct {
	UserID  string `json:"userId"`
	Content string `json:"content"`
}
