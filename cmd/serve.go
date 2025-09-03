package cmd

import (
	"fmt"
	"go-practice-api/middleware"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() // router

	wrappedMux := manager.WrapMux(mux) // wrap the mux with middleware

	initRoutes(mux, manager) // Initialize routes

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", wrappedMux) //start server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
