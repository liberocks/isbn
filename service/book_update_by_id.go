package service

import (
	"context"

	"isbn/dto"
	"isbn/logger"
)

func (s *Service) BookUpdateByID(ctx context.Context, isbn string, book dto.BookUpdateByIDRequest) (dto.BookUpdateByIDResponse, error) {
	// Call the repository method to update a book by ID
	updatedBook, err := s.repo.BookUpdateByID(ctx, isbn, book)
	if err != nil {
		return dto.BookUpdateByIDResponse{}, err
	}

	logger.Logger.Info("Book updated successfully", "ISBN", updatedBook.ISBN, "Title", updatedBook.Title)

	// Map the updated book to the response DTO
	response := dto.BookUpdateByIDResponse{
		ISBN:        updatedBook.ISBN,
		Title:       updatedBook.Title,
		Author:      updatedBook.Author,
		ReleaseDate: updatedBook.ReleaseDate,
	}

	return response, nil
}
