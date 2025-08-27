package cmd

import (
	"fmt"
	"go-practice-api/global_router"
	"go-practice-api/handlers"
	"go-practice-api/middleware"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Logger)

	mux := http.NewServeMux() //router

	mux.Handle("GET /hasan", manager.With(http.HandlerFunc(handlers.Test))) //log

	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{productID}", manager.With(http.HandlerFunc(handlers.GetProductByID)))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", globalRouter) //start server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
