package product

import (
	"encoding/json"
	"fmt"
	"go-practice-api/domain"
	"go-practice-api/utilities"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID) // Use pID to retrieve the product from the database
	if err != nil {
		utilities.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utilities.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	_, err = h.svc.Update(domain.Product{
		ID:          pId,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		utilities.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utilities.SendData(w, http.StatusOK, "Successfully updated product")
}
