package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	_ "server/docs"
	"server/handlers"
)

// @title IWB Submission API
// @version 1.0
// @contact.name   Liam Griffiths
// @contact.email  me@liam-griffiths.co.uk
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

	// Base routes
	router.Get("/api", handlers.Health())
	router.Get("/api/api-docs/*", handlers.SwaggerDocs())
	router.Get("/api/api-docs", handlers.SwaggerDocsRedirect())

	// Products routes
	router.Route("/api/products", func(r chi.Router) {

		r.Get("/", handlers.GetAllProducts) // GET /api/products - all products
		r.Post("/", handlers.CreateProduct) // POST /api/products - new product

		r.Route("/{productID}", func(r chi.Router) {
			r.Get("/", handlers.GetProduct)       // GET /api/products/123
			r.Put("/", handlers.UpdateProduct)    // PUT /api/products/123
			r.Delete("/", handlers.DeleteProduct) // DELETE /api/products/123
		})
	})

	// Serve and log!
	listenOn := ":3000"
	log.Println("Listening on", listenOn)
	log.Fatal(http.ListenAndServe(listenOn, router))
}
