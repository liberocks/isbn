package handler

import (
	"isbn/service"
)

type HandlerInterface interface {
	BookHandlerInterface
}

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
