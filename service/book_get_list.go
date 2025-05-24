package service

import (
	"context"

	"isbn/dto"
	"isbn/logger"
)

func (s *Service) BookGetList(ctx context.Context, book dto.BookGetListQuery) (dto.BookGetListResponse, error) {
	// Call the repository method to get a list of books
	books, count, err := s.repo.BookGetList(ctx, book)
	if err != nil {
		return dto.BookGetListResponse{}, err
	}

	logger.Logger.Info("Book list retrieved successfully", "Count", count, "Page", book.Page, "Limit", book.Limit)

	// Map the books to the response DTO
	response := dto.BookGetListResponse{
		Data:       make([]dto.BookGetListItemResponse, len(books)),
		Total:      count,
		TotalPages: count / book.Limit,
		Limit:      book.Limit,
		Page:       book.Page,
	}

	for i, b := range books {
		response.Data[i] = dto.BookGetListItemResponse{
			ISBN:        b.ISBN,
			Title:       b.Title,
			Author:      b.Author,
			ReleaseDate: b.ReleaseDate,
		}
	}

	return response, nil
}
