package cmd

import (
	"fmt"
	"go-practice-api/config"
	"go-practice-api/middleware"
	"net/http"
	"os"
	"strconv"
)

func Serve() {
	cnf := config.GetConfig()

	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() // router

	wrappedMux := manager.WrapMux(mux) // wrap the mux with middleware

	initRoutes(mux, manager) // Initialize routes

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server is running on port", addr)

	err := http.ListenAndServe(addr, wrappedMux) //start server
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
