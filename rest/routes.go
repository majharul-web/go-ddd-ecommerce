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
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProductList)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct), middleware.AuthenticateJWT))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(handlers.GetProductByID), middleware.AuthenticateJWT))
	mux.Handle("PUT /products/{productId}", manager.With(http.HandlerFunc(handlers.UpdateProduct), middleware.AuthenticateJWT))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(handlers.DeleteProduct), middleware.AuthenticateJWT))

	// ✅ User routes
	mux.Handle("GET /users", manager.With(http.HandlerFunc(handlers.GetUserList)))
	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("GET /users/{userId}", manager.With(http.HandlerFunc(handlers.GetUserByID)))
	mux.Handle("PUT /users/{userId}", manager.With(http.HandlerFunc(handlers.UpdateUser)))
	mux.Handle("DELETE /users/{userId}", manager.With(http.HandlerFunc(handlers.DeleteUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.LoginUser)))
}
