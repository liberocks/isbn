package service_test

import (
	"context"
	"testing"

	"isbn/dto"
	"isbn/repository"
	"isbn/service"
)

func TestBookDeleteByID(t *testing.T) {
	// Initialize the service
	repo := repository.NewRepository()
	svc := service.NewService(repo)

	// Create a new book
	book := dto.BookCreateRequest{
		ISBN:        "978-3-16-148410-0",
		Title:       "Test Book",
		Author:      "John Doe",
		ReleaseDate: "2023-10-01",
	}

	// Call the CreateBook function
	createdBook, err := svc.BookCreate(context.Background(), book)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Delete the book
	err = repo.BookDeleteByID(context.Background(), createdBook.ISBN)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Try to retrieve the deleted book
	_, err = repo.BookGetByID(context.Background(), createdBook.ISBN)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
