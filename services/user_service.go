package services

import (
	"errors"

	"main.go/clients"
	"main.go/dto"
)

type UserService struct {
	client *clients.TcpClient
}

func NewUserService(client *clients.TcpClient) UserService {
	return UserService{client: client}
}

func (s *UserService) GetUsers(payload dto.UserDto) ([]dto.GetUsersResponse, error) {
	client := s.client

	users, err := client.GetUsers(payload)
	if err != nil {
		return []dto.GetUsersResponse{}, errors.New("Erro ao buscar usuários")
	}

	return users, nil
}
