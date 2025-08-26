package cmd

import (
	"fmt"
	"go-practice-api/global_router"
	"go-practice-api/handlers"
	"go-practice-api/middleware"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() //router

	mux.Handle("GET /hasan", middleware.Logger(http.HandlerFunc(handlers.Test))) //log

	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{productID}", middleware.Logger(http.HandlerFunc(handlers.GetProductByID)))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", globalRouter) //start server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
