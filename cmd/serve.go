package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.TestingMiddleware)

	// Initialize the HTTP server
	mux := http.NewServeMux()

	initRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}