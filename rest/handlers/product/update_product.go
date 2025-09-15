package product

import (
	"encoding/json"
	"fmt"
	"go-practice-api/database"
	"go-practice-api/utilities"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID) // Use pID to retrieve the product from the database
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me valid json", 400)
		return
	}

	newProduct.ID = pId

	database.Update(newProduct)

	utilities.SendData(w, "Successfully updated product", 201)
}
