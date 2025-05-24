package repository

import (
	"context"
	"fmt"

	"isbn/dto"
	"isbn/model"
)

func (r *Repository) BookUpdateByID(ctx context.Context, id string, book dto.BookUpdateByIDRequest) (*model.Book, error) {
	// Check if the book exists
	existingBook, exists := bookStore[id]
	if !exists {
		return nil, fmt.Errorf("book with ISBN %s does not exist", id)
	}

	// Update the book details
	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.ReleaseDate = book.ReleaseDate

	bookStore[id] = existingBook

	return &existingBook, nil
}
