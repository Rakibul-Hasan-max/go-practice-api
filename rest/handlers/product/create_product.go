package product

import (
	"encoding/json"
	"fmt"
	"go-practice-api/database"
	"go-practice-api/utilities"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	utilities.SendData(w, createdProduct, 201)
}
