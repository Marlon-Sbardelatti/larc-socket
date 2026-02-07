package routes

import (
	"github.com/go-chi/chi/v5"
	"main.go/clients"
	"main.go/handlers"
	"main.go/services"
)

func RegisterRoutes(r chi.Router) {
	tcpClient := clients.NewTcpClient()
	udpClient := clients.NewUdpClient()

	r.Route("/users", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware)

		userService := services.NewUserService(tcpClient)

		r.Get("/", handlers.GetUsersHandler(userService))
	})

	r.Route("/messages", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware)

		messageService := services.NewMessageService(tcpClient, udpClient)

		r.Get("/", handlers.GetMessagesHandler(messageService))
		r.Post("/", handlers.SendMessageHandler(messageService))
	})

	r.Route("/auth", func(r chi.Router) {
		authService := services.NewAuthService(tcpClient)

		r.Post("/login", handlers.LoginHandler(authService))
	})

}
