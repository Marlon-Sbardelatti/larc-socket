package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"main.go/clients"
	"main.go/dto"
)

type AuthService struct {
	client *clients.TcpClient
}

func NewAuthService(client *clients.TcpClient) AuthService {
	return AuthService{client: client}
}

func (s *AuthService) Login(payload *dto.UserDto) (string, *dto.GetUsersResponse, error) {
	_, err := s.client.GetUsers(dto.UserDto{
		ID:       payload.ID,
		Password: payload.Password,
	})

	// como a requisicao retorna os usuarios ativos, na primeira vez chamad ainda não estamos no resultado
	// de GetUsers, por isso chamamos duas vezes

	users, err := s.client.GetUsers(dto.UserDto{
		ID:       payload.ID,
		Password: payload.Password,
	})

	if err != nil {
		return "", nil, err
	}

	currentUser := GetUserFromUsers(payload.ID, users)
	if currentUser == nil {
		return "", nil, errors.New("Erro ao buscar usuário logado")
	}

	token, err := generateToken(payload.ID, payload.Password)
	if err != nil {
		return "", nil, errors.New("Erro ao criar token")
	}

	return token, currentUser, nil
}

func generateToken(userID int, password string) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"user_id":  userID,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
