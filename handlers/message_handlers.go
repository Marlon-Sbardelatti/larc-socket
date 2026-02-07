package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(payload)
	}
}

func MessagesWSHandler(svc services.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authUser, ok := r.Context().Value("authUser").(dto.UserDto)
		if !ok {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		ticker := time.NewTicker(6 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				messages, err := svc.GetMessages(authUser)
				if err != nil {
					log.Println(err)

					conn.WriteJSON(map[string]string{
						"type":  "error",
						"error": "Erro ao buscar mensagens",
					})
					continue
				}

				if err := conn.WriteJSON(messages); err != nil {
					log.Println(err)
					return
				}

			case <-r.Context().Done():
				return
			}
		}

	}
}
