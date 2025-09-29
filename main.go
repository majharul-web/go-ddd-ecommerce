package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}else{
		fmt.Println("Server started on :8080")
	}
}
