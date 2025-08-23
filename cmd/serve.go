package cmd

import (
	"fmt"
	"go-practice-api/global_router"
	"go-practice-api/handlers"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() //router

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productID}", http.HandlerFunc(handlers.GetProductByID))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", globalRouter) //start server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
