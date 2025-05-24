package repository

import (
	"context"
	"fmt"

	"isbn/dto"
	"isbn/model"
)

func (r *Repository) BookCreate(ctx context.Context, book dto.BookCreateRequest) (*model.Book, error) {
	if _, exists := bookStore[book.ISBN]; exists {
		return nil, fmt.Errorf("book with ISBN %s already exists", book.ISBN)
	}

	newBook := model.Book{
		ISBN:        book.ISBN,
		Title:       book.Title,
		Author:      book.Author,
		ReleaseDate: book.ReleaseDate,
	}

	bookStore[book.ISBN] = newBook

	return &newBook, nil
}
