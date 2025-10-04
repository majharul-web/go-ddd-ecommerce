package database

import "fmt"

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func StoreProduct(p Product) Product {
	productList = append(productList, p)
	return p
}

func GetAllProducts() []Product {
	return productList
}

func GetProductByID(id int) (*Product, error) {
	for _, p := range productList {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("product with ID %d not found", id)
}

func UpdateProduct(id int, updated Product) error {
	for i, p := range productList {
		if p.ID == id {
			updated.ID = id
			productList[i] = updated
			return nil
		}
	}
	return fmt.Errorf("product with ID %d not found", id)
}

func DeleteProduct(id int) error {
	for i, p := range productList {
		if p.ID == id {
			productList = append(productList[:i], productList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("product with ID %d not found", id)
}

func init() {
	fmt.Println("Initializing server...")
	// Initialize some sample products
	productList = []Product{
		{ID: 1, Title: "Product 1", Description: "Description 1", Price: 10.0, ImgUrl: "http://example.com/img1.jpg"},
		{ID: 2, Title: "Product 2", Description: "Description 2", Price: 20.0, ImgUrl: "http://example.com/img2.jpg"},
	}

}
