package repository

import (
	"context"
	"fmt"

	"isbn/dto"
	"isbn/model"
)

func (r *Repository) BookUpdateByID(ctx context.Context, isbn string, book dto.BookUpdateByIDRequest) (*model.Book, error) {
	existingBook, exists := bookStore[isbn]
	if !exists {
		return nil, fmt.Errorf("book with ISBN %s not found", isbn)
	}

	// Update the book fields
	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.ReleaseDate = book.ReleaseDate

	bookStore[isbn] = existingBook

	return &existingBook, nil
}
