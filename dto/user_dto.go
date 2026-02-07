package dto

type UserDto struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type GetUsersResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Wins     int `json:"wins"`
}
