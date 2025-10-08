package repo

import (
	"fmt"
)

// Product struct
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

// ProductRepo interface
type ProductRepo interface {
	Create(p Product) (*Product, error)
	Update(id int, p Product) (*Product, error)
	Delete(id int) error
	Get(id int) (*Product, error)
	List() ([]*Product, error)
}

// productRepo struct
type productRepo struct {
	productList []*Product
}

// NewProductRepo creates a repo and adds initial products
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	repo.generateInitialProducts()
	return repo
}

// Create adds a new product
func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

// Update modifies an existing product by ID
func (r *productRepo) Update(id int, updated Product) (*Product, error) {
	for i, p := range r.productList {
		if p.ID == id {
			updated.ID = id
			r.productList[i] = &updated
			return &updated, nil
		}
	}
	return &updated, fmt.Errorf("product with ID %d not found", id)
}

// Delete removes a product by ID
func (r *productRepo) Delete(id int) error {
	var tempList []*Product
	for _, p := range r.productList {
		if p.ID != id {
			tempList = append(tempList, p)
		}
	}
	if len(tempList) == len(r.productList) {
		return fmt.Errorf("product with ID %d not found", id)
	}
	r.productList = tempList
	return nil
}

// Get retrieves a product by ID
func (r *productRepo) Get(id int) (*Product, error) {
	for _, p := range r.productList {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, fmt.Errorf("product with ID %d not found", id)
}

// List returns all products as pointers
func (r *productRepo) List() (	[]*Product, error) {
	return r.productList, nil
}

// generateInitialProducts adds some sample products
func (r *productRepo) generateInitialProducts() {
	r.productList = []*Product{
		{ID: 1, Title: "Product 1", Description: "Description 1", Price: 10.0, ImgUrl: "http://example.com/img1.jpg"},
		{ID: 2, Title: "Product 2", Description: "Description 2", Price: 20.0, ImgUrl: "http://example.com/img2.jpg"},
	}
}
