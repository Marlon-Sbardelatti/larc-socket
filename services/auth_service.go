package services

import (
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

func (s *AuthService) Login(payload *dto.UserDto) (string, error) {
	_, err := s.client.GetUsers(dto.UserDto{
		ID:       payload.ID,
		Password: payload.Password,
	})

	if err != nil {
		return "", err
	}

	return generateToken(payload.ID, payload.Password)
}

func generateToken(userID string, password string) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"user_id": userID,
		"password": password,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
