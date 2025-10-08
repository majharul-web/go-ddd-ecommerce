package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {
	// Load configuration
	conf := config.GetConfig()

	dbConn, err := db.NewConnection(conf.DB)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbConn)
	userRepo := repo.NewUserRepo(dbConn)

	middlewares := middlewares.NewMiddlewares(conf)

	// Start the REST server
	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(conf, userRepo)

	server := rest.NewServer(conf, productHandler, userHandler)
	server.Start()

}
