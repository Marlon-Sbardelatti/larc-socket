package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"main.go/dto"
	"main.go/services"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(userService services.UserService, messageService services.MessageService) http.HandlerFunc {
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
				users, err := userService.GetUsers(authUser)
				if err != nil {
					log.Println(err)

					conn.WriteJSON(map[string]string{
						"error": "Erro ao buscar usuários",
					})
					continue
				}

				message, err := messageService.GetMessages(authUser)
				if err != nil {
					log.Println(err)

					conn.WriteJSON(map[string]string{
						"error": "Erro ao buscar mensagens",
					})
					continue
				}
				wsMsg := dto.WebsocketMessage{
					Users:   users,
					Message: message,
				}

				if err := conn.WriteJSON(wsMsg); err != nil {
					log.Println(err)
					return
				}

			case <-r.Context().Done():
				return
			}
		}

	}

}
