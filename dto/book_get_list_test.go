package dto_test

import (
	"testing"

	"isbn/dto"
)

func TestBookGetListQuery(t *testing.T) {
	tests := []struct {
		name    string
		query   dto.BookGetListQuery
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid query",
			query: dto.BookGetListQuery{
				Page:  1,
				Limit: 10,
			},
			wantErr: false,
		},
		{
			name: "valid query with max limit",
			query: dto.BookGetListQuery{
				Page:  5,
				Limit: 100,
			},
			wantErr: false,
		},
		{
			name: "valid query with min limit",
			query: dto.BookGetListQuery{
				Page:  1,
				Limit: 1,
			},
			wantErr: false,
		},
		{
			name: "page is zero",
			query: dto.BookGetListQuery{
				Page:  0,
				Limit: 10,
			},
			wantErr: true,
			errMsg:  "page must be greater than 0",
		},
		{
			name: "page is negative",
			query: dto.BookGetListQuery{
				Page:  -1,
				Limit: 10,
			},
			wantErr: true,
			errMsg:  "page must be greater than 0",
		},
		{
			name: "limit is zero",
			query: dto.BookGetListQuery{
				Page:  1,
				Limit: 0,
			},
			wantErr: true,
			errMsg:  "limit must be between 1 and 100",
		},
		{
			name: "limit is negative",
			query: dto.BookGetListQuery{
				Page:  1,
				Limit: -1,
			},
			wantErr: true,
			errMsg:  "limit must be between 1 and 100",
		},
		{
			name: "limit too high",
			query: dto.BookGetListQuery{
				Page:  1,
				Limit: 101,
			},
			wantErr: true,
			errMsg:  "limit must be between 1 and 100",
		},
		{
			name: "both page and limit invalid",
			query: dto.BookGetListQuery{
				Page:  0,
				Limit: 101,
			},
			wantErr: true,
			errMsg:  "page must be greater than 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.query.Validate()
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
