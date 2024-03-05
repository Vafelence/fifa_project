package handlers

import "github.com/Vafelence/fifa_project/internal/services"

type Handler struct {
	service *services.Service
}

func New(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
}
