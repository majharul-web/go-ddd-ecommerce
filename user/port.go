package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

// UserRepo interface
type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	Update(id int, u domain.User) (*domain.User, error)
	Delete(id int) error
	Get(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	List() ([]*domain.User, error)
}