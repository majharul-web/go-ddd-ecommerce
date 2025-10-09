package user

import (
	"ecommerce/domain"
)

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Create(u domain.User) (*domain.User, error) {
	return s.userRepo.Create(u)
}

func (s *service) Update(id int, u domain.User) (*domain.User, error) {
	return s.userRepo.Update(id, u)
}

func (s *service) Delete(id int) error {
	return s.userRepo.Delete(id)
}

func (s *service) Get(id int) (*domain.User, error) {
	return s.userRepo.Get(id)
}

func (s *service) GetByEmail(email string) (*domain.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *service) List() ([]*domain.User, error) {
	return s.userRepo.List()
}