package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //router

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /createProducts", http.HandlerFunc(createProduct))

	fmt.Println("Server is running on port: 8080")

	globalRouter := globalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter) //start server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is a fruit. I love it.",
		Price:       100,
		ImgUrl:      "https://desime.co.uk/cdn/shop/files/orange.jpg?v=1690113703",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is a fruit. I love it.",
		Price:       150,
		ImgUrl:      "https://desime.co.uk/cdn/shop/files/orange.jpg?v=1690113703",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is a fruit. I love it.",
		Price:       200,
		ImgUrl:      "https://desime.co.uk/cdn/shop/files/orange.jpg?v=1690113703",
	}

	productsList = append(productsList, prd1)
	productsList = append(productsList, prd2)
	productsList = append(productsList, prd3)

}
