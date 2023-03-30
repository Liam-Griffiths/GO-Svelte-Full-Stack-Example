package handlers

import (
	"net/http"
	"server/utils"
)

// DeleteProduct
// @Summary Delete product
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/products/{id} [delete]
func DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	utils.Text(writer, 200, "Ya found me!")
}

// UpdateProduct
// @Summary Delete product
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/products/{id} [put]
func UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	utils.Text(writer, 200, "Ya found me!")
}

// CreateProduct
// @Summary Delete product
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/products/{id} [post]
func CreateProduct(writer http.ResponseWriter, request *http.Request) {
	utils.Text(writer, 200, "Ya found me!")
}

// GetProduct
// @Summary Delete product
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/products/{id} [get]
func GetProduct(writer http.ResponseWriter, request *http.Request) {
	utils.Text(writer, 200, "Ya found me! specific")
}

// GetAllProducts
// @Summary Delete product
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/products/{id} [get]
func GetAllProducts(writer http.ResponseWriter, request *http.Request) {
	utils.Text(writer, 200, "Ya found me! all")
}
