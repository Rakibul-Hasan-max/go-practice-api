package product

import (
	"encoding/json"
	"fmt"
	"go-practice-api/repo"
	"go-practice-api/utilities"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utilities.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		utilities.SendError(w, http.StatusInternalServerError, "Internal Server Error")
	}

	utilities.SendData(w, http.StatusCreated, createdProduct)
}
