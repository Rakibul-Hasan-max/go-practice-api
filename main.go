package main

import (
	"fmt"
	"go-practice-api/global_router"
	"go-practice-api/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //router

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /createProducts", http.HandlerFunc(handlers.CreateProduct))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", globalRouter) //start server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
