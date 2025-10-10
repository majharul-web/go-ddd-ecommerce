package product

import (
	"ecommerce/domain"
	productHandler "ecommerce/rest/handlers/product"
)

type Service interface {
	productHandler.Service
}

// ProductRepo interface
type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Update(id int, p domain.Product) (*domain.Product, error)
	Delete(id int) error
	Get(id int) (*domain.Product, error)
	List(page int, limit int) ([]*domain.Product, error)
	Count() (int, error)
}
