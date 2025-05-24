package repository_test

import (
	"context"
	"testing"

	"isbn/dto"
	"isbn/repository"
)

func TestUpdateBookByID(t *testing.T) {
	// Create a new book
	book := dto.CreateBookRequest{
		ISBN:        "978-3-16-148410-0",
		Title:       "Test Book",
		Author:      "John Doe",
		ReleaseDate: "2023-10-01",
	}

	// Call the CreateBook function
	_, err := repository.CreateBook(context.Background(), book)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Update the book
	updateBook := dto.UpdateBookByIDRequest{
		Title:       "Updated Book",
		Author:      "Jane Doe",
		ReleaseDate: "2023-10-02",
	}

	updatedBook, err := repository.UpdateBookByID(context.Background(), book.ISBN, updateBook)
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
	repository.DeleteBookByID(context.Background(), book.ISBN)
}
