package repository

import (
	"context"
	"fmt"
	"isbn/dto"
	"isbn/model"
)

func UpdateBookByID(ctx context.Context, isbn string, book dto.UpdateBookByIDRequest) (*model.Book, error) {
	// Check if the book exists
	existingBook, exists := inMemoryDB[isbn]
	if !exists {
		return nil, fmt.Errorf("book with ISBN %s does not exist", isbn)
	}

	// Update the book details
	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.ReleaseDate = book.ReleaseDate

	inMemoryDB[isbn] = existingBook

	return &existingBook, nil
}
