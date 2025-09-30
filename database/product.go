package database

import "fmt"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var ProductList []Product

func init() {
	fmt.Println("Initializing server...")
	// Initialize some sample products
	ProductList = []Product{
		{ID: 1, Title: "Product 1", Description: "Description 1", Price: 10.0, ImgUrl: "http://example.com/img1.jpg"},
		{ID: 2, Title: "Product 2", Description: "Description 2", Price: 20.0, ImgUrl: "http://example.com/img2.jpg"},
	}

}