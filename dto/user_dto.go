package dto

type UserDto struct {
	ID       int `json:"id"`
	Password string `json:"password"`
}

type GetUsersResponse struct {
	ID       int `json:"id"`
	Username string `json:"username"`
	Wins     int `json:"wins"`
}
