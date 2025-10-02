package cmd

import (
	"fmt"
	"go-practice-api/config"
	"go-practice-api/infra/db"
	"go-practice-api/repo"
	"go-practice-api/rest"
	"go-practice-api/rest/handlers/product"
	"go-practice-api/rest/handlers/review"
	"go-practice-api/rest/handlers/user"
	middleware "go-practice-api/rest/middlewares"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

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
