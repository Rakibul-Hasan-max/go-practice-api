package handlers

import (
	"go-practice-api/database"
	"go-practice-api/utilities"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID) // Use pID to retrieve the product from the database
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	database.Delete(pId)

	utilities.SendData(w, "Successfully deleted product", 200)
}
