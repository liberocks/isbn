package repository

import (
	"context"
	"fmt"

	"isbn/dto"
	"isbn/model"
)

func CreateBook(ctx context.Context, book dto.CreateBookRequest) (*model.Book, error) {
	if _, exists := inMemoryDB[book.ISBN]; exists {
		return nil, fmt.Errorf("book with ISBN %s already exists", book.ISBN)
	}

	newBook := model.Book{
		ISBN:        book.ISBN,
		Title:       book.Title,
		Author:      book.Author,
		ReleaseDate: book.ReleaseDate,
	}

	inMemoryDB[book.ISBN] = newBook

	return &newBook, nil
}
