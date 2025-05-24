package repository

import (
	"context"
	"isbn/model"
	"testing"
)

func TestRepository_BookCount(t *testing.T) {
	tests := []struct {
		name      string
		bookStore map[string]model.Book
		want      int
	}{
		{
			name:      "empty store",
			bookStore: map[string]model.Book{},
			want:      0,
		},
		{
			name: "single book",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Book 1", Author: "Author 1", ReleaseDate: "2020-01-01"},
			},
			want: 1,
		},
		{
			name: "multiple books",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Book 1", Author: "Author 1", ReleaseDate: "2020-01-01"},
				"2": {ISBN: "2", Title: "Book 2", Author: "Author 2", ReleaseDate: "2021-01-01"},
				"3": {ISBN: "3", Title: "Book 3", Author: "Author 1", ReleaseDate: "2022-01-01"},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalBookStore := bookStore
			bookStore = tt.bookStore
			defer func() { bookStore = originalBookStore }()

			r := &Repository{}
			got, err := r.BookCount(context.Background())
			if err != nil {
				t.Errorf("BookCount() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("BookCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_BookAuthorCount(t *testing.T) {
	tests := []struct {
		name      string
		bookStore map[string]model.Book
		want      int
	}{
		{
			name:      "empty store",
			bookStore: map[string]model.Book{},
			want:      0,
		},
		{
			name: "single author",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Book 1", Author: "Author 1", ReleaseDate: "2020-01-01"},
				"2": {ISBN: "2", Title: "Book 2", Author: "Author 1", ReleaseDate: "2021-01-01"},
			},
			want: 1,
		},
		{
			name: "multiple authors",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Book 1", Author: "Author 1", ReleaseDate: "2020-01-01"},
				"2": {ISBN: "2", Title: "Book 2", Author: "Author 2", ReleaseDate: "2021-01-01"},
				"3": {ISBN: "3", Title: "Book 3", Author: "Author 3", ReleaseDate: "2022-01-01"},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalBookStore := bookStore
			bookStore = tt.bookStore
			defer func() { bookStore = originalBookStore }()

			r := &Repository{}
			got, err := r.BookAuthorCount(context.Background())
			if err != nil {
				t.Errorf("BookAuthorCount() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("BookAuthorCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_BookGetOldestReleaseDate(t *testing.T) {
	tests := []struct {
		name      string
		bookStore map[string]model.Book
		want      string
	}{
		{
			name:      "empty store",
			bookStore: map[string]model.Book{},
			want:      "",
		},
		{
			name: "single book",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Book 1", Author: "Author 1", ReleaseDate: "2020-01-01"},
			},
			want: "2020-01-01",
		},
		{
			name: "multiple books",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Book 1", Author: "Author 1", ReleaseDate: "2022-01-01"},
				"2": {ISBN: "2", Title: "Book 2", Author: "Author 2", ReleaseDate: "2020-01-01"},
				"3": {ISBN: "3", Title: "Book 3", Author: "Author 3", ReleaseDate: "2021-01-01"},
			},
			want: "2020-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalBookStore := bookStore
			bookStore = tt.bookStore
			defer func() { bookStore = originalBookStore }()

			r := &Repository{}
			got, err := r.BookGetOldestReleaseDate(context.Background())
			if err != nil {
				t.Errorf("BookGetOldestReleaseDate() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("BookGetOldestReleaseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_BookGetNewestReleaseDate(t *testing.T) {
	tests := []struct {
		name      string
		bookStore map[string]model.Book
		want      string
	}{
		{
			name:      "empty store",
			bookStore: map[string]model.Book{},
			want:      "",
		},
		{
			name: "single book",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Book 1", Author: "Author 1", ReleaseDate: "2020-01-01"},
			},
			want: "2020-01-01",
		},
		{
			name: "multiple books",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Book 1", Author: "Author 1", ReleaseDate: "2020-01-01"},
				"2": {ISBN: "2", Title: "Book 2", Author: "Author 2", ReleaseDate: "2022-01-01"},
				"3": {ISBN: "3", Title: "Book 3", Author: "Author 3", ReleaseDate: "2021-01-01"},
			},
			want: "2022-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalBookStore := bookStore
			bookStore = tt.bookStore
			defer func() { bookStore = originalBookStore }()

			r := &Repository{}
			got, err := r.BookGetNewestReleaseDate(context.Background())
			if err != nil {
				t.Errorf("BookGetNewestReleaseDate() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("BookGetNewestReleaseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_BookGetMostProductiveAuthor(t *testing.T) {
	tests := []struct {
		name      string
		bookStore map[string]model.Book
		want      string
	}{
		{
			name:      "empty store",
			bookStore: map[string]model.Book{},
			want:      "",
		},
		{
			name: "single author",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Book 1", Author: "Author 1", ReleaseDate: "2020-01-01"},
			},
			want: "Author 1",
		},
		{
			name: "most productive author",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Book 1", Author: "Author 1", ReleaseDate: "2020-01-01"},
				"2": {ISBN: "2", Title: "Book 2", Author: "Author 2", ReleaseDate: "2021-01-01"},
				"3": {ISBN: "3", Title: "Book 3", Author: "Author 1", ReleaseDate: "2022-01-01"},
				"4": {ISBN: "4", Title: "Book 4", Author: "Author 1", ReleaseDate: "2023-01-01"},
			},
			want: "Author 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalBookStore := bookStore
			bookStore = tt.bookStore
			defer func() { bookStore = originalBookStore }()

			r := &Repository{}
			got, err := r.BookGetMostProductiveAuthor(context.Background())
			if err != nil {
				t.Errorf("BookGetMostProductiveAuthor() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("BookGetMostProductiveAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_BookGetLongestTitle(t *testing.T) {
	tests := []struct {
		name      string
		bookStore map[string]model.Book
		want      string
	}{
		{
			name:      "empty store",
			bookStore: map[string]model.Book{},
			want:      "",
		},
		{
			name: "single book",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Short", Author: "Author 1", ReleaseDate: "2020-01-01"},
			},
			want: "Short",
		},
		{
			name: "longest title",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Short", Author: "Author 1", ReleaseDate: "2020-01-01"},
				"2": {ISBN: "2", Title: "This is a very long book title", Author: "Author 2", ReleaseDate: "2021-01-01"},
				"3": {ISBN: "3", Title: "Medium length title", Author: "Author 3", ReleaseDate: "2022-01-01"},
			},
			want: "This is a very long book title",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalBookStore := bookStore
			bookStore = tt.bookStore
			defer func() { bookStore = originalBookStore }()

			r := &Repository{}
			got, err := r.BookGetLongestTitle(context.Background())
			if err != nil {
				t.Errorf("BookGetLongestTitle() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("BookGetLongestTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_BookGetShortestTitle(t *testing.T) {
	tests := []struct {
		name      string
		bookStore map[string]model.Book
		want      string
	}{
		{
			name:      "empty store",
			bookStore: map[string]model.Book{},
			want:      "",
		},
		{
			name: "single book",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "Title", Author: "Author 1", ReleaseDate: "2020-01-01"},
			},
			want: "Title",
		},
		{
			name: "shortest title",
			bookStore: map[string]model.Book{
				"1": {ISBN: "1", Title: "This is a very long book title", Author: "Author 1", ReleaseDate: "2020-01-01"},
				"2": {ISBN: "2", Title: "Hi", Author: "Author 2", ReleaseDate: "2021-01-01"},
				"3": {ISBN: "3", Title: "Medium length title", Author: "Author 3", ReleaseDate: "2022-01-01"},
			},
			want: "Hi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalBookStore := bookStore
			bookStore = tt.bookStore
			defer func() { bookStore = originalBookStore }()

			r := &Repository{}
			got, err := r.BookGetShortestTitle(context.Background())
			if err != nil {
				t.Errorf("BookGetShortestTitle() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("BookGetShortestTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
