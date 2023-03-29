package handlers

import (
	"net/http"
	"server/utils"
)

func DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	utils.Text(writer, 200, "Ya found me!")
}

func UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	utils.Text(writer, 200, "Ya found me!")
}

func CreateProduct(writer http.ResponseWriter, request *http.Request) {
	utils.Text(writer, 200, "Ya found me!")
}

func GetProduct(writer http.ResponseWriter, request *http.Request) {
	utils.Text(writer, 200, "Ya found me! specific")
}

func GetAllProducts(writer http.ResponseWriter, request *http.Request) {
	utils.Text(writer, 200, "Ya found me! all")
}
