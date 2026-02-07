package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"main.go/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar .env")
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	routes.RegisterRoutes(r)

	port := ":3000"
	log.Printf("Server running on Port %s", port)
	http.ListenAndServe(port, r)
}
