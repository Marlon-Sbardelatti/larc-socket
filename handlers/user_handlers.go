package handlers

import (
	"encoding/json"
	"net/http"
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



