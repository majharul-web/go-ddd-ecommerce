package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	// manager:= middleware.NewManager()
	
	// Initialize the HTTP server
	mux := http.NewServeMux()

	mux.Handle("GET /test", middleware.Logger(middleware.TestingMiddleware((http.HandlerFunc(handlers.Test)))))

	mux.Handle("GET /products", middleware.Logger(middleware.TestingMiddleware(http.HandlerFunc(handlers.GetProductList))))
	mux.Handle("POST /products", middleware.Logger(middleware.TestingMiddleware(http.HandlerFunc(handlers.CreateProduct))))
	mux.Handle("GET /products/{productId}", middleware.Logger(middleware.TestingMiddleware(http.HandlerFunc(handlers.GetProductByID))))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}