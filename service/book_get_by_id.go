package service

import (
	"context"

	"isbn/dto"
	"isbn/logger"
)

func (s *Service) BookGetByID(ctx context.Context, isbn string) (dto.BookGetByIDResponse, error) {
	// Call the repository method to get a book by ID
	book, err := s.repo.BookGetByID(ctx, isbn)
	if err != nil {
		return dto.BookGetByIDResponse{}, err
	}

	logger.Logger.Info("Book retrieved successfully", "ISBN", book.ISBN, "Title", book.Title)

	// Map the book to the response DTO
	response := dto.BookGetByIDResponse{
		ISBN:        book.ISBN,
		Title:       book.Title,
		Author:      book.Author,
		ReleaseDate: book.ReleaseDate,
	}

	return response, nil
}
