package repository

import (
	"context"
	"isbn/dto"
	"isbn/model"
)

func GetBooks(ctx context.Context, query dto.GetBooksQuery) ([]model.Book, int, error) {
	// Initialize a slice to hold the filtered books
	var filteredBooks []model.Book

	// Query parameters
	limit := query.Limit
	offset := (query.Page - 1) * query.Limit

	// Iterate over the in-memory database and filter books based on the query
	index := 0
	for _, book := range inMemoryDB {
		if index < offset {
			index++
			continue
		}

		if index >= offset+limit {
			break
		}

		filteredBooks = append(filteredBooks, book)
		index++
	}

	// Return the filtered books and the total count
	return filteredBooks, len(inMemoryDB), nil
}
