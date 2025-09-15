package product

import "go-practice-api/rest/middlewares"

type Handler struct {
	middlewares *middlewares.Middlewares
}

func NewHandler(middlewares *middlewares.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}
