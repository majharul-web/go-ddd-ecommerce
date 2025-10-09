package product

import (
	"ecommerce/domain"
)

type Service interface {
	Update(id int, product domain.Product) (*domain.Product, error)
	Create(product domain.Product) (*domain.Product, error)
	Delete(id int) error
	Get(id int) (*domain.Product, error)
	List() ([]*domain.Product, error)
}
