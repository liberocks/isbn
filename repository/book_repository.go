package repository

import (
	"context"

	"isbn/dto"
	"isbn/model"
)

var bookStore = make(map[string]model.Book)

type BookRepositoryInterface interface {
	BookCreate(ctx context.Context, book dto.BookCreateRequest) (*model.Book, error)
	BookGetList(ctx context.Context) ([]model.Book, int, error)
	BookGetByID(ctx context.Context, isbn string) (*model.Book, error)
	BookUpdateByID(ctx context.Context, isbn string, book dto.BookUpdateByIDRequest) (*model.Book, error)
	BookDeleteByID(ctx context.Context, isbn string) error
}
