package product

import (
	"go-practice-api/utilities"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID) // Use pID to retrieve the product from the database
	if err != nil {
		utilities.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	product, err := h.svc.Get(pId)
	if err != nil {
		utilities.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if product == nil {
		utilities.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	utilities.SendData(w, http.StatusOK, product)
}
