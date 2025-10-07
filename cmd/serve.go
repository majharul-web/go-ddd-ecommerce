package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Serve() {
	// Load configuration
	conf := config.GetConfig()

	middlewares := middlewares.NewMiddlewares(conf)

	// Start the REST server
	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()

	server := rest.NewServer(conf, productHandler, userHandler)
	server.Start()

}
