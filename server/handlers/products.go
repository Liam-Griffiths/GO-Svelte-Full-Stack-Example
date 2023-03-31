package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	. "server/models"
	"server/utils"
	"strconv"
)

// GetAllProducts
// @Summary Delete product
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/products/{id} [get]
func GetAllProducts(writer http.ResponseWriter, request *http.Request) {
	products, err := utils.ReadProducts()
	if err != nil {
		utils.Text(writer, 404, "Not found")
		return
	}
	utils.JSON(writer, 200, products)
}

// GetProduct
// @Summary Delete product
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/products/{id} [get]
func GetProduct(writer http.ResponseWriter, request *http.Request) {

	idStr := chi.URLParam(request, "productID")
	fmt.Println(idStr)

	// Convert the id string to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Return an error response if the id is not a valid integer
		utils.Text(writer, 400, "Invalid request")
		return
	}

	product, readErr := utils.ReadProduct(id)
	if readErr != nil {
		utils.Text(writer, 404, "Not found")
		return
	}
	utils.JSON(writer, 200, product)
}

// CreateProduct
// @Summary Delete product
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/products/{id} [post]
func CreateProduct(writer http.ResponseWriter, request *http.Request) {
	// Decode the JSON request body into a Product struct
	var newProduct Product
	err := json.NewDecoder(request.Body).Decode(&newProduct)
	if err != nil {
		// Return an error response if the request body cannot be decoded
		utils.Text(writer, 400, "Invalid request body")
		return
	}

	// Update the product with the new data
	err = utils.CreateProduct(newProduct)
	if err != nil {
		utils.Text(writer, 500, "Something went wrong")
		return
	}

	// Return a success response
	writer.WriteHeader(200)
}

// UpdateProduct
// @Summary Delete product
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/products/{id} [put]
func UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	// Decode the JSON request body into a Product struct
	var newProduct Product
	err := json.NewDecoder(request.Body).Decode(&newProduct)
	if err != nil {
		// Return an error response if the request body cannot be decoded
		utils.Text(writer, 400, "Invalid request body")
		return
	}

	// Update the product with the new data
	err = utils.UpdateProduct(newProduct)
	if err != nil {
		utils.Text(writer, 500, "Something went wrong")
		return
	}

	// Return a success response
	writer.WriteHeader(200)
}

// DeleteProduct
// @Summary Delete product
// @Description  Will respond I'm Alive! if the API is up.
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /api/products/{id} [delete]
func DeleteProduct(writer http.ResponseWriter, request *http.Request) {

	idStr := chi.URLParam(request, "productID")

	// Convert the id string to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Return an error response if the id is not a valid integer
		utils.Text(writer, 400, "Invalid product ID")
		return
	}

	err = utils.DeleteProduct(id)
	if err != nil {
		utils.Text(writer, 500, "Something went wrong")
	}

	writer.WriteHeader(200)
}
