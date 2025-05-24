package service_test

import (
	"context"
	"testing"

	"isbn/dto"
	"isbn/model"
	"isbn/repository"
	"isbn/service"
)

func TestBookAnalyticsGet(t *testing.T) {
	// Initialize the service
	repo := repository.NewRepository()
	svc := service.NewService(repo)

	// Create test books
	testBooks := []dto.BookCreateRequest{
		{
			ISBN:        "978-3-16-148410-0",
			Title:       "Test Book One",
			Author:      "Author A",
			ReleaseDate: "2020-01-01",
		},
		{
			ISBN:        "978-3-16-148410-1",
			Title:       "Test Book Two",
			Author:      "Author A",
			ReleaseDate: "2021-01-01",
		},
		{
			ISBN:        "978-3-16-148410-2",
			Title:       "Another Book",
			Author:      "Author B",
			ReleaseDate: "2022-01-01",
		},
	}

	// Create test books
	for _, book := range testBooks {
		_, err := svc.BookCreate(context.Background(), book)
		if err != nil {
			t.Fatalf("failed to create test book: %v", err)
		}
	}

	// Trigger analytics to populate data
	err := svc.BookAnalyticsTrigger(context.Background())
	if err != nil {
		t.Fatalf("failed to trigger analytics: %v", err)
	}

	// Test BookAnalyticsGet
	analytics, err := svc.BookAnalyticsGet(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify the response structure and basic data
	if analytics.TotalBooks == 0 {
		t.Error("expected TotalBooks to be greater than 0")
	}

	if analytics.TotalAuthors == 0 {
		t.Error("expected TotalAuthors to be greater than 0")
	}

	if analytics.OldestBookReleaseDate == "" {
		t.Error("expected OldestBookReleaseDate to not be empty")
	}

	if analytics.NewestBookReleaseDate == "" {
		t.Error("expected NewestBookReleaseDate to not be empty")
	}

	if analytics.MostProductiveAuthor == "" {
		t.Error("expected MostProductiveAuthor to not be empty")
	}

	if analytics.LongestBookTitle == "" {
		t.Error("expected LongestBookTitle to not be empty")
	}

	if analytics.ShortestBookTitle == "" {
		t.Error("expected ShortestBookTitle to not be empty")
	}

	// Clean up test books
	for _, book := range testBooks {
		repo.BookDeleteByID(context.Background(), book.ISBN)
	}
}

func TestBookAnalyticsGetEmpty(t *testing.T) {
	// Initialize the service with empty repository
	repo := repository.NewRepository()
	svc := service.NewService(repo)

	// Clear any existing analytics data first
	emptyAnalytics := model.BookAnalytics{
		TotalBooks:            0,
		TotalAuthors:          0,
		OldestBookReleaseDate: "",
		NewestBookReleaseDate: "",
		MostProductiveAuthor:  "",
		LongestBookTitle:      "",
		ShortestBookTitle:     "",
	}
	_, err := repo.BookAnalyticsUpdate(context.Background(), emptyAnalytics)
	if err != nil {
		t.Fatalf("failed to clear analytics: %v", err)
	}

	// Test getting analytics when no books exist
	analytics, err := svc.BookAnalyticsGet(context.Background())

	// Should not error even with no data
	if err != nil {
		t.Fatalf("expected no error even with empty data, got %v", err)
	}

	// Verify zero values for empty analytics
	if analytics.TotalBooks != 0 {
		t.Errorf("expected TotalBooks to be 0 for empty database, got %d", analytics.TotalBooks)
	}

	if analytics.TotalAuthors != 0 {
		t.Errorf("expected TotalAuthors to be 0 for empty database, got %d", analytics.TotalAuthors)
	}
}
