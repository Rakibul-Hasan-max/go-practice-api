package handlers

import (
	"go-practice-api/database"
	"go-practice-api/utilities"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productID")

	pID, err := strconv.Atoi(productID) // Use pID to retrieve the product from the database
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	product := database.Get(pID)
	if product == nil {
		utilities.SendError(w, 404, "Product not found")
		return
	}

	utilities.SendData(w, product, 200)
}
