package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me valid json", 400)
		return
	}

	newProduct.ID = len(productsList) + 1

	productsList = append(productsList, newProduct)

	sendData(w, newProduct, 201)
}
