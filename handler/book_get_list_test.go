package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"isbn/dto"
	"isbn/handler"
	"isbn/repository"
	"isbn/service"
)

func TestBookGetList(t *testing.T) {
	// Initialize the dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	tests := []struct {
		name           string
		path           string
		expectedStatus int
		setupFunc      func()
		cleanupFunc    func()
		expectedCount  int
	}{
		{
			name:           "Get list with default pagination",
			path:           "/books",
			expectedStatus: http.StatusOK,
			setupFunc: func() {
				repo.BookCreate(context.Background(), dto.BookCreateRequest{
					ISBN: "1111111111111", Title: "Book 1", Author: "Author 1", ReleaseDate: "2023-01-01",
				})
				repo.BookCreate(context.Background(), dto.BookCreateRequest{
					ISBN: "2222222222222", Title: "Book 2", Author: "Author 2", ReleaseDate: "2023-01-02",
				})
			},
			cleanupFunc: func() {
				repo.BookDeleteByID(context.Background(), "1111111111111")
				repo.BookDeleteByID(context.Background(), "2222222222222")
			},
			expectedCount: 2,
		},
		{
			name:           "Get list with pagination",
			path:           "/books?limit=1&page=1",
			expectedStatus: http.StatusOK,
			setupFunc: func() {
				repo.BookCreate(context.Background(), dto.BookCreateRequest{
					ISBN: "3333333333333", Title: "Book 3", Author: "Author 3", ReleaseDate: "2023-01-03",
				})
				repo.BookCreate(context.Background(), dto.BookCreateRequest{
					ISBN: "4444444444444", Title: "Book 4", Author: "Author 4", ReleaseDate: "2023-01-04",
				})
			},
			cleanupFunc: func() {
				repo.BookDeleteByID(context.Background(), "3333333333333")
				repo.BookDeleteByID(context.Background(), "4444444444444")
			},
			expectedCount: 1,
		},
		{
			name:           "Get empty list",
			path:           "/books",
			expectedStatus: http.StatusOK,
			expectedCount:  0,
		},
		{
			name:           "Invalid path",
			path:           "/invalid",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				tt.setupFunc()
			}

			// Create request
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			// Call the handler
			h.BookGetList(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d. Response body: %s", tt.expectedStatus, w.Code, w.Body.String())
			}

			// For successful retrieval, verify response body
			if tt.expectedStatus == http.StatusOK {
				var response dto.BookGetListResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Fatalf("failed to decode response: %v. Response body: %s", err, w.Body.String())
				}

				if len(response.Data) != tt.expectedCount {
					t.Errorf("expected %d books, got %d", tt.expectedCount, len(response.Data))
				}
			}

			// Cleanup
			if tt.cleanupFunc != nil {
				tt.cleanupFunc()
			}
		})
	}
}
