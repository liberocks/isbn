package repository_test

import (
	"context"
	"testing"

	"isbn/repository"
)

func TestBookAnalyticsGet(t *testing.T) {
	// Initialize the repository
	repo := repository.NewRepository()

	// Call the BookAnalyticsGet function
	analytics, err := repo.BookAnalyticsGet(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check if analytics is not nil
	if analytics == nil {
		t.Fatal("expected analytics to be not nil")
	}

	// Check if analytics has the expected default structure
	if analytics.TotalBooks < 0 {
		t.Errorf("expected TotalBooks to be non-negative, got %d", analytics.TotalBooks)
	}
	if analytics.TotalAuthors < 0 {
		t.Errorf("expected TotalAuthors to be non-negative, got %d", analytics.TotalAuthors)
	}
}
