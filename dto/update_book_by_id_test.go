package dto_test

import (
	"strings"
	"testing"

	"isbn/dto"
)

func TestUpdateBookByIDRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request dto.UpdateBookByIDRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request",
			request: dto.UpdateBookByIDRequest{
				Title:       "Valid Book Title",
				Author:      "Valid Author",
				ReleaseDate: "2023-01-01",
			},
			wantErr: false,
		},
		{
			name: "empty title",
			request: dto.UpdateBookByIDRequest{
				Title:       "",
				Author:      "Valid Author",
				ReleaseDate: "2023-01-01",
			},
			wantErr: true,
			errMsg:  "title is required",
		},
		{
			name: "title too short",
			request: dto.UpdateBookByIDRequest{
				Title:       "AB",
				Author:      "Valid Author",
				ReleaseDate: "2023-01-01",
			},
			wantErr: true,
			errMsg:  "title must be between 3 and 100 characters",
		},
		{
			name: "title too long",
			request: dto.UpdateBookByIDRequest{
				Title:       strings.Repeat("A", 101),
				Author:      "Valid Author",
				ReleaseDate: "2023-01-01",
			},
			wantErr: true,
			errMsg:  "title must be between 3 and 100 characters",
		},
		{
			name: "empty author",
			request: dto.UpdateBookByIDRequest{
				Title:       "Valid Book Title",
				Author:      "",
				ReleaseDate: "2023-01-01",
			},
			wantErr: true,
			errMsg:  "author is required",
		},
		{
			name: "author too short",
			request: dto.UpdateBookByIDRequest{
				Title:       "Valid Book Title",
				Author:      "AB",
				ReleaseDate: "2023-01-01",
			},
			wantErr: true,
			errMsg:  "author must be between 3 and 100 characters",
		},
		{
			name: "author too long",
			request: dto.UpdateBookByIDRequest{
				Title:       "Valid Book Title",
				Author:      strings.Repeat("A", 101),
				ReleaseDate: "2023-01-01",
			},
			wantErr: true,
			errMsg:  "author must be between 3 and 100 characters",
		},
		{
			name: "empty release date",
			request: dto.UpdateBookByIDRequest{
				Title:       "Valid Book Title",
				Author:      "Valid Author",
				ReleaseDate: "",
			},
			wantErr: true,
			errMsg:  "release date is required",
		},
		{
			name: "release date wrong length",
			request: dto.UpdateBookByIDRequest{
				Title:       "Valid Book Title",
				Author:      "Valid Author",
				ReleaseDate: "2023-1-1",
			},
			wantErr: true,
			errMsg:  "release date must be in YYYY-MM-DD format",
		},
		{
			name: "release date wrong format - missing first dash",
			request: dto.UpdateBookByIDRequest{
				Title:       "Valid Book Title",
				Author:      "Valid Author",
				ReleaseDate: "20230101-1",
			},
			wantErr: true,
			errMsg:  "release date must be in YYYY-MM-DD format",
		},
		{
			name: "release date wrong format - missing second dash",
			request: dto.UpdateBookByIDRequest{
				Title:       "Valid Book Title",
				Author:      "Valid Author",
				ReleaseDate: "2023-0101",
			},
			wantErr: true,
			errMsg:  "release date must be in YYYY-MM-DD format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if err.Error() != tt.errMsg {
					t.Errorf("Expected error message '%s', got '%s'", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}
