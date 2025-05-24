package repository

import (
	"context"
	"fmt"
)

func (r *Repository) BookDeleteByID(ctx context.Context, id string) error {
	// Check if the book exists
	if _, exists := bookStore[id]; !exists {
		return fmt.Errorf("book with ISBN %s does not exist", id)
	}

	// Delete the book from the in-memory database
	delete(bookStore, id)

	return nil
}
