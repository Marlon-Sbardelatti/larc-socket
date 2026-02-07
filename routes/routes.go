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

	userService := services.NewUserService(tcpClient)
	messageService := services.NewMessageService(tcpClient, udpClient)
	authService := services.NewAuthService(tcpClient)

	r.Route("/users", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware)


		r.Get("/", handlers.GetUsersHandler(userService))
	})

	r.Route("/messages", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware)


		r.Get("/", handlers.GetMessagesHandler(messageService))
		r.Post("/", handlers.SendMessageHandler(messageService))
	})

	r.Route("/auth", func(r chi.Router) {

		r.Post("/login", handlers.LoginHandler(authService))
	})

	r.Route("/ws", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware)

		r.Get("/users", handlers.UsersWSHandler(userService))
		r.Get("/messages", handlers.MessagesWSHandler(messageService))
	})

}
