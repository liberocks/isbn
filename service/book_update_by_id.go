package service

import (
	"context"

	"isbn/dto"
)

func (s *Service) BookUpdateByID(ctx context.Context, id string, book dto.BookUpdateByIDRequest) (dto.BookUpdateByIDResponse, error) {
	// Call the repository method to update a book by ID
	updatedBook, err := s.repo.BookUpdateByID(ctx, id, book)
	if err != nil {
		return dto.BookUpdateByIDResponse{}, err
	}

	// Map the updated book to the response DTO
	response := dto.BookUpdateByIDResponse{
		ISBN:        updatedBook.ISBN,
		Title:       updatedBook.Title,
		Author:      updatedBook.Author,
		ReleaseDate: updatedBook.ReleaseDate,
	}

	return response, nil
}
