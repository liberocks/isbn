package repository

import (
	"context"
	"fmt"

	"isbn/model"
)

func (r *Repository) BookGetByID(ctx context.Context, id string) (*model.Book, error) {
	// Check if the book exists
	book, exists := bookStore[id]
	if !exists {
		return nil, fmt.Errorf("book with ISBN %s does not exist", id)
	}

	// Return the book details
	return &book, nil
}
