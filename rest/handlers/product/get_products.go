package product

import (
	"go-practice-api/utilities"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.svc.List()
	if err != nil {
		utilities.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utilities.SendData(w, http.StatusOK, productList)
}
