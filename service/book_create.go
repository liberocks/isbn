package service

import (
	"context"

	"isbn/dto"
	"isbn/logger"
)

func (s *Service) BookCreate(ctx context.Context, book dto.BookCreateRequest) (dto.BookCreateResponse, error) {
	// Call the repository method to create a book
	newBook, err := s.repo.BookCreate(ctx, book)
	if err != nil {
		return dto.BookCreateResponse{}, err
	}

	logger.Logger.Info("Book created successfully", "ISBN", newBook.ISBN, "Title", newBook.Title)

	// Map the new book to the response DTO
	response := dto.BookCreateResponse{
		ISBN:        newBook.ISBN,
		Title:       newBook.Title,
		Author:      newBook.Author,
		ReleaseDate: newBook.ReleaseDate,
	}

	return response, nil
}
