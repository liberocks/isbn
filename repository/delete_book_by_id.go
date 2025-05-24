package repository

import (
	"context"
	"fmt"
)

func DeleteBookByID(ctx context.Context, isbn string) error {
	// Check if the book exists
	if _, exists := inMemoryDB[isbn]; !exists {
		return fmt.Errorf("book with ISBN %s does not exist", isbn)
	}

	// Delete the book from the in-memory database
	delete(inMemoryDB, isbn)

	return nil
}
