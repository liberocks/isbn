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

func NewBookHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
