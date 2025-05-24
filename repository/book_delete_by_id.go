package repository

import (
	"context"
	"fmt"
)

func (r *Repository) BookDeleteByID(ctx context.Context, isbn string) error {
	if _, exists := bookStore[isbn]; !exists {
		return fmt.Errorf("book with ISBN %s not found", isbn)
	}

	delete(bookStore, isbn)
	return nil
}
