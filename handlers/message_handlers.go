package handlers

import (
	"encoding/json"
	"net/http"

	"main.go/dto"
	"main.go/services"
)

func GetMessagesHandler(svc services.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authUser, ok := r.Context().Value("authUser").(dto.UserDto)
		if !ok {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}

		response, err := svc.GetMessages(authUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func SendMessageHandler(svc services.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authUser, ok := r.Context().Value("authUser").(dto.UserDto)
		if !ok {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}

		var payload dto.SendMessageRequest

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Mensagem inválida", http.StatusBadRequest)
			return
		}

		payload.Sender = &authUser

		svc.SendMessage(payload)

		res := dto.SendMessageResponse{
			SenderId:   payload.Sender.ID,
			ReceiverID: payload.ReceiverID,
			Content:    payload.Content,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
