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

func TestBookGetByID(t *testing.T) {
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
	}{
		{
			name:           "Valid book retrieval",
			path:           "/books/1234567890123",
			expectedStatus: http.StatusOK,
			setupFunc: func() {
				repo.BookCreate(context.Background(), dto.BookCreateRequest{
					ISBN:        "1234567890123",
					Title:       "Test Book",
					Author:      "John Doe",
					ReleaseDate: "2023-01-01",
				})
			},
			cleanupFunc: func() {
				repo.BookDeleteByID(context.Background(), "1234567890123")
			},
		},
		{
			name:           "Book not found",
			path:           "/books/9999999999999",
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Missing book ID",
			path:           "/books/",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid path",
			path:           "/invalid/123",
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
			h.BookGetByID(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d. Response body: %s", tt.expectedStatus, w.Code, w.Body.String())
			}

			// For successful retrieval, verify response body
			if tt.expectedStatus == http.StatusOK {
				var response dto.BookGetByIDResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Fatalf("failed to decode response: %v. Response body: %s", err, w.Body.String())
				}

				if response.ISBN != "1234567890123" {
					t.Errorf("expected ISBN 1234567890123, got %s", response.ISBN)
				}
				if response.Title != "Test Book" {
					t.Errorf("expected Title 'Test Book', got %s", response.Title)
				}
			}

			// Cleanup
			if tt.cleanupFunc != nil {
				tt.cleanupFunc()
			}
		})
	}
}
