package product

import (
	"fmt"
	"go-practice-api/utilities"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID) // Use pID to retrieve the product from the database
	if err != nil {
		fmt.Println(err)
		utilities.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.svc.Delete(pId)
	if err != nil {
		fmt.Println(err)
		utilities.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utilities.SendData(w, http.StatusOK, "Successfully deleted product")
}
