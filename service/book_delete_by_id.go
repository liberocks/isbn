package service

import (
	"context"

	"isbn/dto"
)

func (s *Service) BookDeleteByID(ctx context.Context, id string) (dto.BookDeleteByIDResponse, error) {
	// Call the repository method to delete a book
	err := s.repo.BookDeleteByID(ctx, id)
	if err != nil {
		return dto.BookDeleteByIDResponse{}, err
	}

	// Map the response DTO
	response := dto.BookDeleteByIDResponse{
		Message: "Book deleted successfully",
	}

	return response, nil
}
