package rest

import (
	"fmt"
	"go-practice-api/config"
	"go-practice-api/rest/handlers/product"
	"go-practice-api/rest/handlers/user"
	middleware "go-practice-api/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start(cnf config.Config) {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() // router

	wrappedMux := manager.WrapMux(mux) // wrap the mux with middleware

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server is running on port", addr)

	err := http.ListenAndServe(addr, wrappedMux) //start server
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
