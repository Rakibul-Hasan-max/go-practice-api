package cmd

import (
	"go-practice-api/config"
	"go-practice-api/repo"
	"go-practice-api/rest"
	"go-practice-api/rest/handlers/product"
	"go-practice-api/rest/handlers/review"
	"go-practice-api/rest/handlers/user"
	middleware "go-practice-api/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)
	server.Start()
}
