package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(conf config.Config) {
	manager := middlewares.NewManager()
	manager.Use(middlewares.Preflight, middlewares.Cors, middlewares.Logger)

	mux := http.NewServeMux()
	wrappedMux := manager.WrappedMux(mux)

	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(conf.HttpPort)
	fmt.Println("Starting server on", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}

}
