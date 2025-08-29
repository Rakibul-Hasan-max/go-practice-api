package cmd

import (
	"fmt"
	"go-practice-api/middleware"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Logger)

	mux := http.NewServeMux() // router

	initRoutes(mux, manager) // Initialize routes

	globalRouter := middleware.CorsWithPreflight(mux)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", globalRouter) //start server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
