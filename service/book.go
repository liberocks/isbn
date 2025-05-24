package service

import (
	"context"

	"isbn/dto"
)

type BookServiceInterface interface {
	BookCreate(ctx context.Context, book dto.BookCreateRequest) (dto.BookCreateResponse, error)
	BookGetList(ctx context.Context, query dto.BookGetListQuery) (dto.BookGetListResponse, error)
	BookGetByID(ctx context.Context, isbn string) (dto.BookGetByIDResponse, error)
	BookUpdateByID(ctx context.Context, isbn string, book dto.BookUpdateByIDRequest) (dto.BookUpdateByIDResponse, error)
	BookDeleteByID(ctx context.Context, isbn string) (dto.BookDeleteByIDResponse, error)
	TriggerBookAnalysis(ctx context.Context, isbn string) error
	GetBookAnalytics(ctx context.Context) (dto.BookAnalyticsGetResponse, error)
}
