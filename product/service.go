package product

import (
	"ecommerce/domain"
)

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{productRepo: productRepo}
}

func (s *service) Create(product domain.Product) (*domain.Product, error) {
	return s.productRepo.Create(product)
}

func (s *service) Update(id int, product domain.Product) (*domain.Product, error) {
	return s.productRepo.Update(id, product)
}

func (s *service) Delete(id int) error {
	return s.productRepo.Delete(id)
}

func (s *service) Get(id int) (*domain.Product, error) {
	return s.productRepo.Get(id)
}

func (s *service) List() ([]*domain.Product, error) {
	return s.productRepo.List()
}
