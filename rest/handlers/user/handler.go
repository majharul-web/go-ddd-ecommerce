package user

import (
	"ecommerce/repo"
	"ecommerce/config"
)

type Handler struct {
	// Add any dependencies like services or repositories here
	conf *config.Config
	userRepo repo.UserRepo
}

func NewHandler(conf *config.Config, userRepo repo.UserRepo) *Handler {
	return &Handler{
		conf: conf,
		userRepo: userRepo,
	}
}
