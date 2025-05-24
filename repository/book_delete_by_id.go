package repository

import (
	"context"
	"fmt"
)

func (r *Repository) BookDeleteByID(ctx context.Context, isbn string) error {
	// Check if the book exists
	if _, exists := bookStore[isbn]; !exists {
		return fmt.Errorf("book with ISBN %s does not exist", isbn)
	}

	// Delete the book from the in-memory database
	delete(bookStore, isbn)

	return nil
}
