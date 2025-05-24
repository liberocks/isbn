package repository_test

import (
	"context"
	"testing"

	"isbn/dto"
	"isbn/repository"
)

func TestDeleteBookByID(t *testing.T) {
	// Create a new book
	book := dto.CreateBookRequest{
		ISBN:        "978-3-16-148410-0",
		Title:       "Test Book",
		Author:      "John Doe",
		ReleaseDate: "2023-10-01",
	}

	// Call the CreateBook function
	createdBook, err := repository.CreateBook(context.Background(), book)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Delete the book
	err = repository.DeleteBookByID(context.Background(), createdBook.ISBN)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Try to retrieve the deleted book
	_, err = repository.GetBookByID(context.Background(), createdBook.ISBN)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
