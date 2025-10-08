package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Serve() {
	// Load configuration
	conf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middlewares.NewMiddlewares(conf)

	// Start the REST server
	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(conf,userRepo)

	server := rest.NewServer(conf, productHandler, userHandler)
	server.Start()

}
