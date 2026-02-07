package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"main.go/dto"
	"main.go/services"
)

func LoginHandler(svc services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var payload dto.UserDto
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Mensagem inválida", http.StatusBadRequest)
			return
		}

		token, currentUser, err := svc.Login(&payload)
		if err != nil {
			http.Error(w, "Usuário inválido", http.StatusUnauthorized)
			return
		}

		cookie := &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   3600,
		}

		http.SetCookie(w, cookie)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(currentUser)
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value

		var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)
		password := claims["password"].(string)

		authUser := dto.UserDto{
			ID:       userID,
			Password: password,
		}

		ctx := context.WithValue(r.Context(), "authUser", authUser)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
