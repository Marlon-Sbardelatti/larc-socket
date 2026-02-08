package clients

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"main.go/dto"
)

type TcpClient struct {
	address string
}

func NewTcpClient() *TcpClient {
	larcAddres := os.Getenv("LARC_ADDRESS")
	tcpPort := os.Getenv("TCP_PORT")
	address := fmt.Sprintf("%s:%s", larcAddres, tcpPort)

	return &TcpClient{
		address: address,
	}

}

func (c *TcpClient) GetUsers(payload dto.UserDto) ([]dto.GetUsersResponse, error) {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return []dto.GetUsersResponse{}, err
	}
	defer conn.Close()

	cmd := fmt.Sprintf("GET USERS %d:%s\n", payload.ID, payload.Password)
	_, err = conn.Write([]byte(cmd))
	if err != nil {
		return []dto.GetUsersResponse{}, err
	}

	reader := bufio.NewReader(conn)
	resp, err := reader.ReadString('\n')
	if err != nil {
		return []dto.GetUsersResponse{}, err

	}

	users, err := parseUsers(resp)
	if err != nil {
		return []dto.GetUsersResponse{}, err

	}

	return users, nil
}

func parseUsers(res string) ([]dto.GetUsersResponse, error) {
	parts := strings.Split(res, ":")

	var users []dto.GetUsersResponse

	for i := 0; i < len(parts)-1; i += 3 {
		wins, err := strconv.Atoi(parts[i+2])
		if err != nil {
			return nil, errors.New("Erro ao converter usuário")
		}

		id, err := strconv.Atoi(parts[i])
		if err != nil {
			return []dto.GetUsersResponse{}, err
		}

		user := dto.GetUsersResponse{
			ID:       id,
			Username: parts[i+1],
			Wins:     wins,
		}

		users = append(users, user)
	}

	return users, nil
}

func (c *TcpClient) GetMessage(payload dto.UserDto) (string, error) {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cmd := fmt.Sprintf("GET MESSAGE %d:%s\n", payload.ID, payload.Password)
	_, err = conn.Write([]byte(cmd))
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(conn)
	resp, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return resp, nil
}
