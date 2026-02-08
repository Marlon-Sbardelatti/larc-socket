package dto

type WebsocketMessage struct {
	Users   []GetUsersResponse `json:"users"`
	Message GetMessageResponse `json:"message"`
}
