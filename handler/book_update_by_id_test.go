package handler_test

import (
	"bytes"
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

func TestBookUpdateByID(t *testing.T) {
	// Initialize the dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	tests := []struct {
		name           string
		path           string
		requestBody    interface{}
		expectedStatus int
		setupFunc      func()
		cleanupFunc    func()
	}{
		{
			name: "Valid book update",
			path: "/books/1234567890123",
			requestBody: dto.BookUpdateByIDRequest{
				Title:       "Updated Book",
				Author:      "Updated Author",
				ReleaseDate: "2023-12-01",
			},
			expectedStatus: http.StatusOK,
			setupFunc: func() {
				repo.BookCreate(context.Background(), dto.BookCreateRequest{
					ISBN:        "1234567890123",
					Title:       "Original Book",
					Author:      "Original Author",
					ReleaseDate: "2023-01-01",
				})
			},
			cleanupFunc: func() {
				repo.BookDeleteByID(context.Background(), "1234567890123")
			},
		},
		{
			name: "Book not found",
			path: "/books/9999999999999",
			requestBody: dto.BookUpdateByIDRequest{
				Title:       "Updated Book",
				Author:      "Updated Author",
				ReleaseDate: "2023-12-01",
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Invalid JSON",
			path:           "/books/1234567890123",
			requestBody:    "invalid json",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Missing book ID",
			path:           "/books/",
			requestBody:    dto.BookUpdateByIDRequest{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid path",
			path:           "/invalid/123",
			requestBody:    dto.BookUpdateByIDRequest{},
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				tt.setupFunc()
			}

			// Prepare request body
			var reqBody []byte
			if str, ok := tt.requestBody.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, _ = json.Marshal(tt.requestBody)
			}

			// Create request
			req := httptest.NewRequest(http.MethodPut, tt.path, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Call the handler
			h.BookUpdateByID(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d. Response body: %s", tt.expectedStatus, w.Code, w.Body.String())
			}

			// For successful update, verify response body
			if tt.expectedStatus == http.StatusOK {
				var response dto.BookUpdateByIDResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Fatalf("failed to decode response: %v. Response body: %s", err, w.Body.String())
				}

				expectedBook := tt.requestBody.(dto.BookUpdateByIDRequest)
				if response.Title != expectedBook.Title {
					t.Errorf("expected Title %s, got %s", expectedBook.Title, response.Title)
				}
				if response.Author != expectedBook.Author {
					t.Errorf("expected Author %s, got %s", expectedBook.Author, response.Author)
				}
				if response.ReleaseDate != expectedBook.ReleaseDate {
					t.Errorf("expected ReleaseDate %s, got %s", expectedBook.ReleaseDate, response.ReleaseDate)
				}
			}

			// Cleanup
			if tt.cleanupFunc != nil {
				tt.cleanupFunc()
			}
		})
	}
}
