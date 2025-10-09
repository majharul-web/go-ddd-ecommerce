package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"ecommerce/user"
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

	err = db.MigrateDB(dbConn, "./migrations")
	if err != nil {
		fmt.Println("Failed to apply migrations:", err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbConn)
	userRepo := repo.NewUserRepo(dbConn)

	// domains
	userService := user.NewService(userRepo)

	middlewares := middlewares.NewMiddlewares(conf)

	// Start the REST server
	productHandler := productHandler.NewHandler(middlewares, productRepo)
	userHandler := userHandler.NewHandler(conf, userService)

	server := rest.NewServer(conf, productHandler, userHandler)
	server.Start()

}
