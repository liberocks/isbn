package repository_test

import (
	"context"
	"testing"

	"isbn/dto"
	"isbn/repository"
)

func TestBookUpdateByID(t *testing.T) {
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
	_, err := repo.BookCreate(context.Background(), book)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Update the book
	updateBook := dto.BookUpdateByIDRequest{
		Title:       "Updated Book",
		Author:      "Jane Doe",
		ReleaseDate: "2023-10-02",
	}

	updatedBook, err := repo.BookUpdateByID(context.Background(), book.ISBN, updateBook)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Check if the book was updated successfully
	if updatedBook.ISBN != book.ISBN {
		t.Errorf("expected ISBN %s, got %s", book.ISBN, updatedBook.ISBN)
	}
	if updatedBook.Title != updateBook.Title {
		t.Errorf("expected Title %s, got %s", updateBook.Title, updatedBook.Title)
	}
	if updatedBook.Author != updateBook.Author {
		t.Errorf("expected Author %s, got %s", updateBook.Author, updatedBook.Author)
	}
	if updatedBook.ReleaseDate != updateBook.ReleaseDate {
		t.Errorf("expected ReleaseDate %s, got %s", updateBook.ReleaseDate, updatedBook.ReleaseDate)
	}

	// Clean up
	repo.BookDeleteByID(context.Background(), book.ISBN)
}
