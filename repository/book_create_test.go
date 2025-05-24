package repository_test

import (
	"context"
	"testing"

	"isbn/dto"
	"isbn/repository"
)

func TestBookCreate(t *testing.T) {
	// Initialize the repository
	repo := repository.NewRepository()

	// Create a new book
	book := dto.BookCreateRequest{
		ISBN:        "978-3-16-148410-0",
		Title:       "Test Book",
		Author:      "John Doe",
		ReleaseDate: "2023-10-01",
	}

	// Call the CreateBook function
	createdBook, err := repo.BookCreate(context.Background(), book)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check if the book was created successfully
	if createdBook.ISBN != book.ISBN {
		t.Errorf("expected ISBN %s, got %s", book.ISBN, createdBook.ISBN)
	}
	if createdBook.Title != book.Title {
		t.Errorf("expected Title %s, got %s", book.Title, createdBook.Title)
	}
	if createdBook.Author != book.Author {
		t.Errorf("expected Author %s, got %s", book.Author, createdBook.Author)
	}
	if createdBook.ReleaseDate != book.ReleaseDate {
		t.Errorf("expected ReleaseDate %s, got %s", book.ReleaseDate, createdBook.ReleaseDate)
	}

	// Clean up
	repo.BookDeleteByID(context.Background(), createdBook.ISBN)
}
