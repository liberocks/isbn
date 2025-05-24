package service

import (
	"context"

	"isbn/dto"
)

func (s *Service) BookGetByID(ctx context.Context, id string) (dto.BookGetByIDResponse, error) {
	// Call the repository method to get a book by ID
	book, err := s.repo.BookGetByID(ctx, id)
	if err != nil {
		return dto.BookGetByIDResponse{}, err
	}

	// Map the book to the response DTO
	response := dto.BookGetByIDResponse{
		ISBN:        book.ISBN,
		Title:       book.Title,
		Author:      book.Author,
		ReleaseDate: book.ReleaseDate,
	}

	return response, nil
}
