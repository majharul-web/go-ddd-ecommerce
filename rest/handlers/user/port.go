package user

import (
	"ecommerce/domain"
)

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Get(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(id int, user domain.User) (*domain.User, error)
	Delete(id int) error
	List() ([]*domain.User, error)
}