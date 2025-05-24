package service

import (
	"context"

	"isbn/dto"
	"isbn/logger"
)

func (s *Service) BookAnalyticsGet(ctx context.Context) (dto.BookAnalyticsGetResponse, error) {
	// Call the repository method to get book analytics
	analytics, err := s.repo.BookAnalyticsGet(ctx)
	if err != nil {
		return dto.BookAnalyticsGetResponse{}, err
	}

	logger.Logger.Info("Book analytics retrieved successfully")

	// Map the response DTO
	response := dto.BookAnalyticsGetResponse{
		TotalBooks:            analytics.TotalBooks,
		TotalAuthors:          analytics.TotalAuthors,
		OldestBookReleaseDate: analytics.OldestBookReleaseDate,
		NewestBookReleaseDate: analytics.NewestBookReleaseDate,
		MostProductiveAuthor:  analytics.MostProductiveAuthor,
		LongestBookTitle:      analytics.LongestBookTitle,
		ShortestBookTitle:     analytics.ShortestBookTitle,
	}

	return response, nil
}
