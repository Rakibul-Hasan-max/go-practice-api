package handlers

import (
	"go-practice-api/product"
	"go-practice-api/utilities"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utilities.SendData(w, product.ProductsList, 200)
}
