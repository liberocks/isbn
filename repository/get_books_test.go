package repository_test

import (
	"context"
	"testing"

	"isbn/dto"
	"isbn/repository"
)

func TestGetBooksID(t *testing.T) {
	// Create multiple books
	books := []dto.CreateBookRequest{
		{
			ISBN:        "978-3-16-148410-0",
			Title:       "Test Book 1",
			Author:      "John Doe",
			ReleaseDate: "2023-10-01",
		},
		{
			ISBN:        "978-3-16-148410-1",
			Title:       "Test Book 2",
			Author:      "Mary Jane",
			ReleaseDate: "2023-10-02",
		},
		{
			ISBN:        "978-3-16-148410-2",
			Title:       "Test Book 3",
			Author:      "Jane Smith",
			ReleaseDate: "2023-10-03",
		},
	}

	for _, book := range books {
		// Call the CreateBook function
		_, err := repository.CreateBook(context.Background(), book)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	// Retrieve all books with page 2 limit 2
	query := dto.GetBooksQuery{
		Page:  2,
		Limit: 2,
	}
	retreivedBooks, total, err := repository.GetBooks(context.Background(), query)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(retreivedBooks) != 1 {
		t.Errorf("expected 1 books, got %d", len(retreivedBooks))
	}

	if total != 3 {
		t.Errorf("expected total 3 books, got %d", total)
	}

	// Clean up
	for _, book := range books {
		repository.DeleteBookByID(context.Background(), book.ISBN)
	}
}
