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

func TestBookCreate(t *testing.T) {
	// Initialize the dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		setupFunc      func()
		cleanupFunc    func()
	}{
		{
			name: "Valid book creation",
			requestBody: dto.BookCreateRequest{
				ISBN:        "1234567890123",
				Title:       "Test Book",
				Author:      "John Doe",
				ReleaseDate: "2023-10-01",
			},
			expectedStatus: http.StatusCreated,
			cleanupFunc: func() {
				repo.BookDeleteByID(context.Background(), "1234567890123")
			},
		},
		{
			name:           "Invalid JSON",
			requestBody:    "invalid json",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid request - empty ISBN",
			requestBody: dto.BookCreateRequest{
				ISBN:        "",
				Title:       "Test Book",
				Author:      "John Doe",
				ReleaseDate: "2023-10-01",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Duplicate ISBN",
			requestBody: dto.BookCreateRequest{
				ISBN:        "1234567890123",
				Title:       "Test Book 2",
				Author:      "Jane Doe",
				ReleaseDate: "2023-10-02",
			},
			expectedStatus: http.StatusInternalServerError,
			setupFunc: func() {
				repo.BookCreate(context.Background(), dto.BookCreateRequest{
					ISBN:        "1234567890123",
					Title:       "Existing Book",
					Author:      "Existing Author",
					ReleaseDate: "2023-01-01",
				})
			},
			cleanupFunc: func() {
				repo.BookDeleteByID(context.Background(), "1234567890123")
			},
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
			req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Call the handler
			h.BookCreate(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d. Response body: %s", tt.expectedStatus, w.Code, w.Body.String())
			}

			// For successful creation, verify response body
			if tt.expectedStatus == http.StatusCreated {
				var response dto.BookCreateResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Fatalf("failed to decode response: %v. Response body: %s", err, w.Body.String())
				}

				expectedBook := tt.requestBody.(dto.BookCreateRequest)
				if response.ISBN != expectedBook.ISBN {
					t.Errorf("expected ISBN %s, got %s", expectedBook.ISBN, response.ISBN)
				}
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

func TestBookCreateInvalidPath(t *testing.T) {
	// Initialize the dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	// Create request with invalid path
	req := httptest.NewRequest(http.MethodPost, "/invalid", nil)
	w := httptest.NewRecorder()

	// Call the handler
	h.BookCreate(w, req)

	// Check status code
	if w.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}
