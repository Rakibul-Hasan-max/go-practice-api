package cmd

import (
	"go-practice-api/handlers"
	"go-practice-api/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	//log
	mux.Handle(
		"GET /hasan",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)

	mux.Handle(
		"GET /products/{productID}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)

}
