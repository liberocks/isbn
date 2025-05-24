package repository

import (
	"context"

	"isbn/dto"
	"isbn/model"
)

var bookStore = make(map[string]model.Book)
var bookAnalytics = model.BookAnalytics{
	TotalBooks:            0,
	TotalAuthors:          0,
	OldestBookReleaseDate: "",
	NewestBookReleaseDate: "",
	MostProductiveAuthor:  "",
	LongestBookTitle:      "",
	ShortestBookTitle:     "",
}

type BookRepositoryInterface interface {
	BookCreate(ctx context.Context, book dto.BookCreateRequest) (*model.Book, error)
	BookGetList(ctx context.Context) ([]model.Book, int, error)
	BookGetByID(ctx context.Context, isbn string) (*model.Book, error)
	BookUpdateByID(ctx context.Context, isbn string, book dto.BookUpdateByIDRequest) (*model.Book, error)
	BookDeleteByID(ctx context.Context, isbn string) error

	// Book Analytics
	BookAnalyticsUpdate(ctx context.Context, analytics model.BookAnalytics) (*model.BookAnalytics, error)
	BookAnalyticsGet(ctx context.Context) (*model.BookAnalytics, error)
	BookCount(ctx context.Context) (int, error)
	BookAuthorCount(ctx context.Context) (int, error)
	BookGetOldestReleaseDate(ctx context.Context) (string, error)
	BookGetNewestReleaseDate(ctx context.Context) (string, error)
	BookGetMostProductiveAuthor(ctx context.Context) (string, error)
	BookGetLongestTitle(ctx context.Context) (string, error)
	BookGetShortestTitle(ctx context.Context) (string, error)
}
