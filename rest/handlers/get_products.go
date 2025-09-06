package handlers

import (
	"go-practice-api/database"
	"go-practice-api/utilities"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utilities.SendData(w, database.List(), 200)
}
