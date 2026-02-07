package clients

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"main.go/dto"
	"main.go/models"
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

func (c *TcpClient) GetUsers(user *models.User) (string, error) {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cmd := fmt.Sprintf("GET USERS %s:%s\n", user.ID, user.Password)
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

func (c *TcpClient) GetMessage(req *dto.GetMessageRequest) (string, error) {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cmd := fmt.Sprintf("GET MESSAGE %s:%s\n", req.User.ID, req.User.Password)
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
