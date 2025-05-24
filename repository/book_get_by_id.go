package repository

import (
	"context"
	"fmt"

	"isbn/model"
)

func (r *Repository) BookGetByID(ctx context.Context, isbn string) (*model.Book, error) {
	book, exists := bookStore[isbn]
	if !exists {
		return nil, fmt.Errorf("book with ISBN %s not found", isbn)
	}

	return &book, nil
}
