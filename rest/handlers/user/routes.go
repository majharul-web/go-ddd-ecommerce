package user

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /users", manager.With(http.HandlerFunc(h.GetUserList)))
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("GET /users/{userId}", manager.With(http.HandlerFunc(h.GetUserByID)))
	mux.Handle("PUT /users/{userId}", manager.With(http.HandlerFunc(h.UpdateUser)))
	mux.Handle("DELETE /users/{userId}", manager.With(http.HandlerFunc(h.DeleteUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.LoginUser)))
}
