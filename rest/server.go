package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
	conf           *config.Config
}

func NewServer(conf *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		conf:           conf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (s *Server) Start() {
	manager := middlewares.NewManager()
	manager.Use(middlewares.Preflight, middlewares.Cors, middlewares.Logger)

	mux := http.NewServeMux()
	wrappedMux := manager.WrappedMux(mux)

	s.productHandler.RegisterRoutes(mux, manager)
	s.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(s.conf.HttpPort)
	fmt.Println("Starting server on", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}

}
