package repository

import "isbn/model"

var inMemoryDB = make(map[string]model.Book)

type BookRepository struct {
}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}
