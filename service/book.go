package service

import (
	"context"

	"isbn/dto"
)

type BookServiceInterface interface {
	BookCreate(ctx context.Context, book dto.BookCreateRequest) (dto.BookCreateResponse, error)
	BookGetList(ctx context.Context) (dto.BookGetListResponse, error)
	BookGetByID(ctx context.Context, id string) (dto.BookGetByIDResponse, error)
	BookUpdateByID(ctx context.Context, id string, book dto.BookUpdateByIDRequest) (dto.BookUpdateByIDResponse, error)
	BookDeleteByID(ctx context.Context, id string) (dto.BookDeleteByIDResponse, error)
}
