package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"go-practice-api/product"
	"go-practice-api/utilities"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me valid json", 400)
		return
	}

	newProduct.ID = len(product.ProductsList) + 1

	product.ProductsList = append(product.ProductsList, newProduct)

	utilities.SendData(w, newProduct, 201)
}
