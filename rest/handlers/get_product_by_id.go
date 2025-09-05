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

	for _, product := range database.ProductsList {
		if product.ID == pID {
			utilities.SendData(w, product, 200)
			return
		}
	}
	utilities.SendData(w, "Data not found", 404)
}
