package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
	"main.go/dto"
	"main.go/services"
)

func GetUsersHandler(svc services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authUser, ok := r.Context().Value("authUser").(dto.UserDto)
		if !ok {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}

		users, err := svc.GetUsers(authUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}



var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func UsersWSHandler(svc services.UserService) http.HandlerFunc {
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
				users, err := svc.GetUsers(authUser)
				if err != nil {
					log.Println(err)

					conn.WriteJSON(map[string]string{
						"type":  "error",
						"error": "Erro ao buscar usuários",
					})
					continue
				}

				if err := conn.WriteJSON(users); err != nil {
					log.Println(err)
					return
				}

			case <-r.Context().Done():
				return
			}
		}

	}
}
