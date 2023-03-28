package main

import (
	"log"
	"net/http"
	responses "server/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	log.Println("Starting Server...")
	// Create our router
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		responses.Text(w, 200, "Alive!")
	})

	// Serve!
	port := ":3000"
	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
