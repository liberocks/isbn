package repository

import (
	"context"
	"fmt"

	"isbn/dto"
	"isbn/model"
)

func (r *Repository) BookUpdateByID(ctx context.Context, isbn string, book dto.BookUpdateByIDRequest) (*model.Book, error) {
	// Check if the book exists
	existingBook, exists := bookStore[isbn]
	if !exists {
		return nil, fmt.Errorf("book with ISBN %s does not exist", isbn)
	}

	// Update the book details
	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.ReleaseDate = book.ReleaseDate

	bookStore[isbn] = existingBook

	return &existingBook, nil
}
