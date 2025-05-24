package repository_test

import (
	"context"
	"testing"

	"isbn/dto"
	"isbn/repository"
)

func TestBookGetList(t *testing.T) {
	// Initialize the repository
	repo := repository.NewRepository()

	// Create multiple books
	books := []dto.BookCreateRequest{
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
		_, err := repo.BookCreate(context.Background(), book)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	// Retrieve all books with page 2 limit 2
	query := dto.BookGetListQuery{
		Page:  2,
		Limit: 2,
	}
	retreivedBooks, total, err := repo.BookGetList(context.Background(), query)
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
		repo.BookDeleteByID(context.Background(), book.ISBN)
	}
}
