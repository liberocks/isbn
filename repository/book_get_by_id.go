package repository

import (
	"context"
	"fmt"

	"isbn/model"
)

func (r *Repository) BookGetByID(ctx context.Context, isbn string) (*model.Book, error) {
	// Check if the book exists
	book, exists := bookStore[isbn]
	if !exists {
		return nil, fmt.Errorf("book with ISBN %s does not exist", isbn)
	}

	// Return the book details
	return &book, nil
}
