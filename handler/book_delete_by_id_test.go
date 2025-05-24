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

func TestBookDeleteByID(t *testing.T) {
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
			name:           "Valid book deletion",
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
		{
			name:           "Empty book ID",
			path:           "/books/",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				tt.setupFunc()
			}

			// Create request
			req := httptest.NewRequest(http.MethodDelete, tt.path, nil)

			// Create response recorder
			w := httptest.NewRecorder()

			// Call the handler
			h.BookDeleteByID(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d. Response body: %s", tt.expectedStatus, w.Code, w.Body.String())
			}

			// For successful deletion, verify response body
			if tt.expectedStatus == http.StatusOK {
				var response dto.BookDeleteByIDResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Fatalf("failed to decode response: %v. Response body: %s", err, w.Body.String())
				}
			}

			// Cleanup
			if tt.cleanupFunc != nil {
				tt.cleanupFunc()
			}
		})
	}
}

func TestBookDeleteByIDWithExistingBook(t *testing.T) {
	// Initialize the dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	// Create a book first
	testISBN := "9876543210987"
	_, err := repo.BookCreate(context.Background(), dto.BookCreateRequest{
		ISBN:        testISBN,
		Title:       "Book to Delete",
		Author:      "Test Author",
		ReleaseDate: "2023-05-01",
	})
	if err != nil {
		t.Fatalf("failed to create test book: %v", err)
	}

	// Create delete request
	req := httptest.NewRequest(http.MethodDelete, "/books/"+testISBN, nil)
	w := httptest.NewRecorder()

	// Call the handler
	h.BookDeleteByID(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d. Response body: %s", http.StatusOK, w.Code, w.Body.String())
	}

	// Verify the book was actually deleted by trying to get it
	_, err = repo.BookGetByID(context.Background(), testISBN)
	if err == nil {
		t.Error("expected book to be deleted, but it still exists")
	}
}
