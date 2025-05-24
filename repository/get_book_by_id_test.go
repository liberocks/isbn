package repository_test

import (
	"context"
	"testing"

	"isbn/dto"
	"isbn/repository"
)

func TestGetBookByID(t *testing.T) { // Create a new book
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

	// Check if the book is stored in the in-memory database
	retreivedBook, err := repository.GetBookByID(context.Background(), book.ISBN)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if retreivedBook.ISBN != book.ISBN {
		t.Errorf("expected ISBN %s, got %s", book.ISBN, retreivedBook.ISBN)
	}
	if retreivedBook.Title != book.Title {
		t.Errorf("expected Title %s, got %s", book.Title, retreivedBook.Title)
	}
	if retreivedBook.Author != book.Author {
		t.Errorf("expected Author %s, got %s", book.Author, retreivedBook.Author)
	}
	if retreivedBook.ReleaseDate != book.ReleaseDate {
		t.Errorf("expected ReleaseDate %s, got %s", book.ReleaseDate, retreivedBook.ReleaseDate)
	}

	// Clean up
	repository.DeleteBookByID(context.Background(), book.ISBN)
}
