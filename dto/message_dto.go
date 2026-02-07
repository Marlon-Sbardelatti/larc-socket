package dto

type SendMessageRequest struct {
	Sender     *UserDto `json:"sender"`
	ReceiverID string   `json:"receiverId"`
	Content    string   `json:"content"`
}

type GetMessageResponse struct {
	UserID  string
	Content string
}
