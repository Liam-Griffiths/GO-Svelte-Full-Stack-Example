package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"server/handlers"
	"server/utils"
)

func main() {
	log.Println("Starting Server...")

	// Create our router
	router := chi.NewRouter()

	// Setup our middlewares
	router.Use(middleware.Logger) // Logs API interactions
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Alive check.
	router.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		utils.Text(w, 200, "Alive!")
	})

	// Swagger docs.
	router.Get("/api/api-docs", func(w http.ResponseWriter, r *http.Request) {
		utils.Text(w, 200, "Swagger goes here!")
	})

	// Products routes
	router.Route("/api/products", func(r chi.Router) {

		r.Get("/", handlers.GetAllProducts) // GET /products - all products
		r.Post("/", handlers.CreateProduct) // POST /products - new product

		r.Route("/{productID}", func(r chi.Router) {
			r.Get("/", handlers.GetProduct)       // GET /products/123
			r.Put("/", handlers.UpdateProduct)    // PUT /products/123
			r.Delete("/", handlers.DeleteProduct) // DELETE /products/123
		})
	})

	// Serve and log!
	listenOn := ":3001"
	log.Println("Listening on", listenOn)
	log.Fatal(http.ListenAndServe(listenOn, router))
}
