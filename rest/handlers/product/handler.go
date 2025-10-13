package product

import (
	"go-practice-api/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	svc         Service
}

func NewHandler(
	middlewares *middlewares.Middlewares,
	svc Service,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc:         svc,
	}
}
