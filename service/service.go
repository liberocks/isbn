package service

import (
	"isbn/repository"
)

type BookService struct {
	repo *repository.Repository
}

func NewBookService(repo *repository.Repository) *BookService {
	return &BookService{repo: repo}
}
