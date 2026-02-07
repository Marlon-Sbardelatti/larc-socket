package services

import (
	"errors"
	"strings"

	"main.go/clients"
	"main.go/dto"
)

type MessageService struct {
	tcpClient *clients.TcpClient
	udpClient *clients.UdpClient
}

func NewMessageService(tcpClient *clients.TcpClient, udpClient *clients.UdpClient) MessageService {
	return MessageService{tcpClient: tcpClient, udpClient: udpClient}
}

func (s *MessageService) GetMessages(payload dto.UserDto) (dto.GetMessageResponse, error) {

	client := s.tcpClient
	res, err := client.GetMessage(payload)
	if err != nil {
		return dto.GetMessageResponse{}, errors.New("Erro ao enviar mensagem")
	}

	content := strings.Split(res, ":")

	response := dto.GetMessageResponse{
		UserID:  content[0],
		Content: content[1],
	}

	return response, nil
}

func (s *MessageService) SendMessage(payload dto.SendMessageRequest) error {
	client := s.udpClient

	err := client.SendMessageUDP(payload)
	if err != nil {
		return errors.New("Erro ao enviar mensagem")
	}

	return nil
}
