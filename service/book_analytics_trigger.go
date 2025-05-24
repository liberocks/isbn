package service

import (
	"context"
	"fmt"
	"sync"

	"isbn/logger"
	"isbn/model"
)

func (s *Service) BookAnalyticsTrigger(ctx context.Context) error {
	var wg sync.WaitGroup
	var analytics model.BookAnalytics

	// Result channels for each goroutine
	var totalBooks int
	var totalAuthors int
	var oldestReleaseDate string
	var newestReleaseDate string
	var mostProductiveAuthor string
	var longestTitle string
	var shortestTitle string

	// Error tracking
	var hasErrors bool
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		count, err := s.bookCount(ctx)
		if err != nil {
			logger.Logger.Error("Failed to get book count", "error", err)
			mu.Lock()
			hasErrors = true
			mu.Unlock()
			return
		}
		totalBooks = count
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		count, err := s.bookAuthorCount(ctx)
		if err != nil {
			logger.Logger.Error("Failed to get author count", "error", err)
			mu.Lock()
			hasErrors = true
			mu.Unlock()
			return
		}
		totalAuthors = count
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		oldest, err := s.bookGetOldestReleaseDate(ctx)
		if err != nil {
			logger.Logger.Error("Failed to get oldest book release date", "error", err)
			mu.Lock()
			hasErrors = true
			mu.Unlock()
			return
		}
		oldestReleaseDate = oldest
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		newest, err := s.bookGetNewestReleaseDate(ctx)
		if err != nil {
			logger.Logger.Error("Failed to get newest book release date", "error", err)
			mu.Lock()
			hasErrors = true
			mu.Unlock()
			return
		}
		newestReleaseDate = newest
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		author, err := s.bookGetMostProductiveAuthor(ctx)
		if err != nil {
			logger.Logger.Error("Failed to get most productive author", "error", err)
			mu.Lock()
			hasErrors = true
			mu.Unlock()
			return
		}
		mostProductiveAuthor = author
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		title, err := s.bookGetLongestTitle(ctx)
		if err != nil {
			logger.Logger.Error("Failed to get longest book title", "error", err)
			mu.Lock()
			hasErrors = true
			mu.Unlock()
			return
		}
		longestTitle = title
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		title, err := s.bookGetShortestTitle(ctx)
		if err != nil {
			logger.Logger.Error("Failed to get shortest book title", "error", err)
			mu.Lock()
			hasErrors = true
			mu.Unlock()
			return
		}
		shortestTitle = title
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Assign results to analytics struct
	analytics.TotalBooks = totalBooks
	analytics.TotalAuthors = totalAuthors
	analytics.OldestBookReleaseDate = oldestReleaseDate
	analytics.NewestBookReleaseDate = newestReleaseDate
	analytics.MostProductiveAuthor = mostProductiveAuthor
	analytics.LongestBookTitle = longestTitle
	analytics.ShortestBookTitle = shortestTitle

	fmt.Printf("Analytics: %+v\n", analytics)

	if hasErrors {
		logger.Logger.Warn("Some analytics operations failed, proceeding with partial data")
	}

	fmt.Printf("Analytics: %+v\n", analytics)

	// Update the analytics in the repository
	_, err := s.repo.BookAnalyticsUpdate(ctx, analytics)
	if err != nil {
		logger.Logger.Error("Failed to update book analytics", "error", err)
		return err
	}

	return nil
}

func (s *Service) bookCount(ctx context.Context) (int, error) {
	// Call the repository method to count books
	count, err := s.repo.BookCount(ctx)
	if err != nil {
		return 0, err
	}

	logger.Logger.Info("Book count retrieved successfully", "Count", count)

	return count, nil
}

func (s *Service) bookAuthorCount(ctx context.Context) (int, error) {
	// Call the repository method to count authors
	count, err := s.repo.BookAuthorCount(ctx)
	if err != nil {
		return 0, err
	}

	logger.Logger.Info("Author count retrieved successfully", "Count", count)

	return count, nil
}

func (s *Service) bookGetOldestReleaseDate(ctx context.Context) (string, error) {
	// Call the repository method to get the oldest book release date
	oldest, err := s.repo.BookGetOldestReleaseDate(ctx)
	if err != nil {
		return "", err
	}

	logger.Logger.Info("Oldest book release date retrieved successfully", "OldestBookReleaseDate", oldest)

	return oldest, nil
}

func (s *Service) bookGetNewestReleaseDate(ctx context.Context) (string, error) {
	// Call the repository method to get the newest book release date
	newest, err := s.repo.BookGetNewestReleaseDate(ctx)
	if err != nil {
		return "", err
	}

	logger.Logger.Info("Newest book release date retrieved successfully", "NewestBookReleaseDate", newest)

	return newest, nil
}

func (s *Service) bookGetMostProductiveAuthor(ctx context.Context) (string, error) {
	// Call the repository method to get the most productive author
	author, err := s.repo.BookGetMostProductiveAuthor(ctx)
	if err != nil {
		return "", err
	}

	logger.Logger.Info("Most productive author retrieved successfully", "Author", author)

	return author, nil
}

func (s *Service) bookGetLongestTitle(ctx context.Context) (string, error) {
	// Call the repository method to get the longest book title
	title, err := s.repo.BookGetLongestTitle(ctx)
	if err != nil {
		return "", err
	}

	logger.Logger.Info("Longest book title retrieved successfully", "LongestBookTitle", title)

	return title, nil
}

func (s *Service) bookGetShortestTitle(ctx context.Context) (string, error) {
	// Call the repository method to get the shortest book title
	title, err := s.repo.BookGetShortestTitle(ctx)
	if err != nil {
		return "", err
	}

	logger.Logger.Info("Shortest book title retrieved successfully", "ShortestBookTitle", title)

	return title, nil
}
