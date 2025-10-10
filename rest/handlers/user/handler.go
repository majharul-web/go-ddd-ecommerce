package user

import (
	"ecommerce/config"
)

type Handler struct {
	// Add any dependencies like services or repositories here
	conf *config.Config
	svc Service
}

func NewHandler(conf *config.Config, svc Service) *Handler {
	return &Handler{
		conf: conf,
		svc:  svc,
	}
}
