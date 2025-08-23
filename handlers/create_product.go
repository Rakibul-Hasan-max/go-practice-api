package handlers

import (
	"encoding/json"
	"fmt"
	"go-practice-api/database"
	"go-practice-api/utilities"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me valid json", 400)
		return
	}

	newProduct.ID = len(database.ProductsList) + 1

	database.ProductsList = append(database.ProductsList, newProduct)

	utilities.SendData(w, newProduct, 201)
}
