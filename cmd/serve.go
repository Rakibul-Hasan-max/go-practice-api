package cmd

import (
	"go-practice-api/config"
	"go-practice-api/rest"
	"go-practice-api/rest/handlers/product"
	"go-practice-api/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(productHandler, userHandler)
	server.Start(cnf)
}
