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

}
