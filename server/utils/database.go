package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	. "server/models"
)

func ReadProducts() ([]Product, error) {
	out := []Product{}
	err := loadData(&out)

	if err != nil {
		fmt.Println(err)
		return out, errors.New("Not Found")
	}

	return out, nil
}

func ReadProduct(id int) (Product, error) {
	products, _ := ReadProducts()

	for i := range products {
		if products[i].ProductID == id {
			return products[i], nil
		}
	}

	return Product{}, errors.New("Not Found")
}

func CreateProduct(newProduct Product) error {
	products, _ := ReadProducts()

	// Get the ID of the last product in the slice
	lastID := 0
	if len(products) > 0 {
		lastID = products[len(products)-1].ProductID
	}
	newProduct.ProductID = lastID + 1

	newProducts := append(products, newProduct)
	err := writeData(newProducts)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(updatedProduct Product) error {
	products, _ := ReadProducts()
	for i := range products {
		if products[i].ProductID == updatedProduct.ProductID {
			products[i] = updatedProduct
			break
		}
	}

	err := writeData(products)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id int) error {
	products, _ := ReadProducts()

	newProducts := []Product{}
	for i, product := range products {
		if product.ProductID == id {
			newProducts = append(products[:i], products[i+1:]...)
		}
	}

	err := writeData(newProducts)
	if err != nil {
		return err
	}

	return nil
}

func loadData(data any) error {
	// Open the JSON file and read its contents
	jsonFile, err := os.Open("db/db.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into the provided interface
	err = json.Unmarshal(byteValue, data)
	if err != nil {
		return err
	}

	return nil
}

func writeData(data any) error {
	// Marshal the data into a JSON byte slice
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Write the JSON byte slice to the file
	err = os.WriteFile("db/db.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
