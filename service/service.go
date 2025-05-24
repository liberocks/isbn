package service

import (
	"isbn/repository"
)

type ServiceInterface interface {
	BookServiceInterface
}

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}
