package routes

import (
	"github.com/go-chi/chi/v5"
	"main.go/handlers"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.GetUsers())
	})

	r.Route("/messages", func(r chi.Router) {
		r.Get("/", handlers.GetMessages())
		r.Post("/", handlers.SendMessage())
	})

}
