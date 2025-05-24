package service_test

import (
	"context"
	"testing"

	"isbn/dto"
	"isbn/repository"
	"isbn/service"
)

func TestBookAnalyticsTrigger(t *testing.T) {
	// Initialize the service
	repo := repository.NewRepository()
	svc := service.NewService(repo)

	// Create test books to generate analytics data
	testBooks := []dto.BookCreateRequest{
		{
			ISBN:        "978-3-16-148410-0",
			Title:       "A Short Title",
			Author:      "John Doe",
			ReleaseDate: "2020-01-01",
		},
		{
			ISBN:        "978-3-16-148410-1",
			Title:       "A Very Long Book Title That Should Be The Longest",
			Author:      "John Doe",
			ReleaseDate: "2023-12-31",
		},
		{
			ISBN:        "978-3-16-148410-2",
			Title:       "X",
			Author:      "Jane Smith",
			ReleaseDate: "2022-06-15",
		},
	}

	// Create test books
	for _, book := range testBooks {
		_, err := svc.BookCreate(context.Background(), book)
		if err != nil {
			t.Fatalf("failed to create test book: %v", err)
		}
	}

	// Trigger analytics calculation
	err := svc.BookAnalyticsTrigger(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify analytics were calculated by retrieving them
	analytics, err := svc.BookAnalyticsGet(context.Background())
	if err != nil {
		t.Fatalf("failed to get analytics after trigger: %v", err)
	}

	// Verify basic analytics data
	if analytics.TotalBooks != 3 {
		t.Errorf("expected TotalBooks to be 3, got %d", analytics.TotalBooks)
	}

	if analytics.TotalAuthors != 2 {
		t.Errorf("expected TotalAuthors to be 2, got %d", analytics.TotalAuthors)
	}

	// Clean up test books
	for _, book := range testBooks {
		repo.BookDeleteByID(context.Background(), book.ISBN)
	}
}
