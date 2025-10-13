package cmd

import (
	"fmt"
	"go-practice-api/config"
	"go-practice-api/infra/db"
	"go-practice-api/product"
	"go-practice-api/repo"
	"go-practice-api/rest"
	prdcthandler "go-practice-api/rest/handlers/product"
	"go-practice-api/rest/handlers/review"
	usrHandler "go-practice-api/rest/handlers/user"
	middleware "go-practice-api/rest/middlewares"
	"go-practice-api/user"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := prdcthandler.NewHandler(middlewares, prdctSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)
	server.Start()
}
