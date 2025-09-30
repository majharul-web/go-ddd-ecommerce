package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

type Product struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProductList(w http.ResponseWriter, r *http.Request) {
	handleCors(w, r)
	handlePreflight(w, r)

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w, r)
	handlePreflight(w, r)

	if r.Method != "POST" {
		http.Error(w, "Please use POST method", 400)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	sendData(w, newProduct,201)
}

func handleCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Habib")
	w.Header().Set("Content-Type", "application/json")
	
}

func handlePreflight(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	}
}

func sendData(w http.ResponseWriter, data interface{},statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
func main() {
	
	// Initialize some sample products
	productList = []Product{
		{ID: 1, Title: "Product 1", Description: "Description 1", Price: 10.0, ImgUrl: "http://example.com/img1.jpg"},
		{ID: 2, Title: "Product 2", Description: "Description 2", Price: 20.0, ImgUrl: "http://example.com/img2.jpg"},
	}


	// Initialize the HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/products", getProductList)
	mux.HandleFunc("/products/create", createProduct)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	fmt.Println("Initializing server...")
}