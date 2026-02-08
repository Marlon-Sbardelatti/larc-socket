package clients

import (
	"fmt"
	"main.go/dto"
	"net"
	"os"
)

type UdpClient struct {
	address string
}

func NewUdpClient() *UdpClient {
	larcAddres := os.Getenv("LARC_ADDRESS")
	udpPort := os.Getenv("UDP_PORT")
	address := fmt.Sprintf("%s:%s", larcAddres, udpPort)

	return &UdpClient{
		address: address,
	}
}

func (c *UdpClient) SendMessageUDP(req dto.SendMessageRequest) error {
	conn, err := net.Dial("udp", c.address)
	if err != nil {
		return err
	}
	defer conn.Close()

	cmd := fmt.Sprintf(
		"SEND MESSAGE %d:%s:%d:%s",
		req.Sender.ID,
		req.Sender.Password,
		req.ReceiverID,
		req.Content,
	)

	_, err = conn.Write([]byte(cmd))
	return err
}
