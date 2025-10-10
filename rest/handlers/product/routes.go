package product

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProductList)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct), h.middlewares.AuthenticateJWT))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(h.GetProductByID), h.middlewares.AuthenticateJWT))
	mux.Handle("PUT /products/{productId}", manager.With(http.HandlerFunc(h.UpdateProduct), h.middlewares.AuthenticateJWT))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(h.DeleteProduct), h.middlewares.AuthenticateJWT))

}
