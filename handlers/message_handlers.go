package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"main.go/clients"
	"main.go/dto"
)

func GetMessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload dto.GetMessageRequest

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Mensagem inválida", http.StatusBadRequest)
			return
		}

		client := clients.NewTcpClient()
		res, err := client.GetMessage(&payload)
		if err != nil {
			http.Error(w, "Erro ao enviar mensagem", http.StatusServiceUnavailable)
		}

		content := strings.Split(res, ":")

		response := dto.GetMessageResponse{UserID: content[0], Content: content[1]}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func SendMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload dto.SendMessageRequest

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Mensagem inválida", http.StatusBadRequest)
			return
		}

		client := clients.NewUdpClient()

		err = client.SendMessageUDP(&payload)
		if err != nil {
			http.Error(w, "Erro ao enviar mensagem", http.StatusServiceUnavailable)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(payload)
	}
}
