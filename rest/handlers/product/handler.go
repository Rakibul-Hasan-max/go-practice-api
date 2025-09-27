package product

import (
	"go-practice-api/rest/middlewares"
	"go-practice-api/repo"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middlewares.Middlewares,
	productRepo repo.ProductRepo,
	) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
