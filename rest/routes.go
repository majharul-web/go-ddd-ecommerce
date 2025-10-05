package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// ✅ Test route
	mux.Handle("GET /test",
		manager.With(http.HandlerFunc(handlers.Test)))

	// ✅ Product routes
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProductList), middleware.MoreMiddleware))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle("PUT /products/{productId}", manager.With(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(handlers.DeleteProduct)))

	// ✅ User routes
	mux.Handle("GET /users", manager.With(http.HandlerFunc(handlers.GetUserList), middleware.MoreMiddleware))
	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("GET /users/{userId}", manager.With(http.HandlerFunc(handlers.GetUserByID)))
	mux.Handle("PUT /users/{userId}", manager.With(http.HandlerFunc(handlers.UpdateUser)))
	mux.Handle("DELETE /users/{userId}", manager.With(http.HandlerFunc(handlers.DeleteUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.LoginUser)))
}
