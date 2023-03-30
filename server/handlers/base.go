package handlers

import (
	"github.com/swaggo/http-swagger"
	"net/http"
	"server/utils"
)

// Health
// @Summary Health Check
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api [get]
func Health() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Text(w, 200, "Alive!")
	}
}

// SwaggerDocs
// @Summary Swagger docs.
// @Description  Swagger docs.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/api-docs [get]
func SwaggerDocs() http.HandlerFunc {
	return httpSwagger.WrapHandler
}

func SwaggerDocsRedirect() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Redirect(w, "/api/api-docs/")
	}
}
