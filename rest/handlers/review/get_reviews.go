package review

import (
	"go-practice-api/database"
	"go-practice-api/utilities"
	"net/http"
)

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	utilities.SendData(w, database.ListReviews(), 200)
}
