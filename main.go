package main

import (
	"fmt"
	"go-practice-api/global_router"
	"go-practice-api/handlers"
	"go-practice-api/product"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //router

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /createProducts", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("Server is running on port: 8080")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter) //start server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	prd1 := product.Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is a fruit. I love it.",
		Price:       100,
		ImgUrl:      "https://desime.co.uk/cdn/shop/files/orange.jpg?v=1690113703",
	}

	prd2 := product.Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is a fruit. I love it.",
		Price:       150,
		ImgUrl:      "https://desime.co.uk/cdn/shop/files/orange.jpg?v=1690113703",
	}

	prd3 := product.Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is a fruit. I love it.",
		Price:       200,
		ImgUrl:      "https://desime.co.uk/cdn/shop/files/orange.jpg?v=1690113703",
	}

	product.ProductsList = append(product.ProductsList, prd1)
	product.ProductsList = append(product.ProductsList, prd2)
	product.ProductsList = append(product.ProductsList, prd3)

}
