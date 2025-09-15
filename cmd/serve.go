package cmd

import (
	"go-practice-api/config"
	"go-practice-api/rest"
	"go-practice-api/rest/handlers/product"
	"go-practice-api/rest/handlers/review"
	"go-practice-api/rest/handlers/user"
	middleware "go-practice-api/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)
	server.Start()
}
