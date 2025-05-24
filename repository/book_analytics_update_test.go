package repository_test

import (
	"context"
	"testing"

	"isbn/model"
	"isbn/repository"
)

func TestBookAnalyticsUpdate(t *testing.T) {
	// Initialize the repository
	repo := repository.NewRepository()

	// Create test analytics data
	testAnalytics := model.BookAnalytics{
		TotalBooks:            10,
		TotalAuthors:          5,
		OldestBookReleaseDate: "2000-01-01",
		NewestBookReleaseDate: "2023-12-31",
		MostProductiveAuthor:  "Jane Smith",
		LongestBookTitle:      "A Very Long Book Title That Goes On And On",
		ShortestBookTitle:     "Short",
	}

	// Call the BookAnalyticsUpdate function
	updatedAnalytics, err := repo.BookAnalyticsUpdate(context.Background(), testAnalytics)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check if the returned analytics is not nil
	if updatedAnalytics == nil {
		t.Fatal("expected updated analytics to be not nil")
	}

	// Verify that the analytics were updated correctly
	if updatedAnalytics.TotalBooks != testAnalytics.TotalBooks {
		t.Errorf("expected TotalBooks %d, got %d", testAnalytics.TotalBooks, updatedAnalytics.TotalBooks)
	}
	if updatedAnalytics.TotalAuthors != testAnalytics.TotalAuthors {
		t.Errorf("expected TotalAuthors %d, got %d", testAnalytics.TotalAuthors, updatedAnalytics.TotalAuthors)
	}
	if updatedAnalytics.OldestBookReleaseDate != testAnalytics.OldestBookReleaseDate {
		t.Errorf("expected OldestBookReleaseDate %s, got %s", testAnalytics.OldestBookReleaseDate, updatedAnalytics.OldestBookReleaseDate)
	}
	if updatedAnalytics.NewestBookReleaseDate != testAnalytics.NewestBookReleaseDate {
		t.Errorf("expected NewestBookReleaseDate %s, got %s", testAnalytics.NewestBookReleaseDate, updatedAnalytics.NewestBookReleaseDate)
	}
	if updatedAnalytics.MostProductiveAuthor != testAnalytics.MostProductiveAuthor {
		t.Errorf("expected MostProductiveAuthor %s, got %s", testAnalytics.MostProductiveAuthor, updatedAnalytics.MostProductiveAuthor)
	}
	if updatedAnalytics.LongestBookTitle != testAnalytics.LongestBookTitle {
		t.Errorf("expected LongestBookTitle %s, got %s", testAnalytics.LongestBookTitle, updatedAnalytics.LongestBookTitle)
	}
	if updatedAnalytics.ShortestBookTitle != testAnalytics.ShortestBookTitle {
		t.Errorf("expected ShortestBookTitle %s, got %s", testAnalytics.ShortestBookTitle, updatedAnalytics.ShortestBookTitle)
	}

	// Verify that subsequent calls to BookAnalyticsGet return the updated data
	retrievedAnalytics, err := repo.BookAnalyticsGet(context.Background())
	if err != nil {
		t.Fatalf("expected no error when retrieving analytics, got %v", err)
	}

	if retrievedAnalytics.TotalBooks != testAnalytics.TotalBooks {
		t.Errorf("expected retrieved TotalBooks %d, got %d", testAnalytics.TotalBooks, retrievedAnalytics.TotalBooks)
	}
	if retrievedAnalytics.MostProductiveAuthor != testAnalytics.MostProductiveAuthor {
		t.Errorf("expected retrieved MostProductiveAuthor %s, got %s", testAnalytics.MostProductiveAuthor, retrievedAnalytics.MostProductiveAuthor)
	}
}
