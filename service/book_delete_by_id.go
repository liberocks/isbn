package service

import (
	"context"
	"log/slog"

	"isbn/dto"
)

func (s *Service) BookDeleteByID(ctx context.Context, isbn string) (dto.BookDeleteByIDResponse, error) {
	// Call the repository method to delete a book
	err := s.repo.BookDeleteByID(ctx, isbn)
	if err != nil {
		return dto.BookDeleteByIDResponse{}, err
	}

	slog.Info("Book deleted successfully", "ISBN", isbn)

	// Map the response DTO
	response := dto.BookDeleteByIDResponse{
		Message: "Book deleted successfully",
	}

	return response, nil
}
