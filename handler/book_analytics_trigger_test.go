package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"isbn/dto"
	"isbn/handler"
	"isbn/repository"
	"isbn/service"
)

func TestBookAnalyticsTrigger(t *testing.T) {
	// Initialize the dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	tests := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Valid analytics trigger",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request
			req := httptest.NewRequest(http.MethodPost, "/analytics", nil)

			// Create response recorder
			w := httptest.NewRecorder()

			// Call the handler
			h.BookAnalyticsTrigger(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d. Response body: %s", tt.expectedStatus, w.Code, w.Body.String())
			}

			// For successful requests, verify response structure
			if tt.expectedStatus == http.StatusOK {
				var response dto.BookAnalyticsTriggerResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Fatalf("failed to decode response: %v. Response body: %s", err, w.Body.String())
				}

				// Verify response message
				expectedMessage := "Book analytics triggered successfully"
				if response.Message != expectedMessage {
					t.Errorf("expected message %s, got %s", expectedMessage, response.Message)
				}

				time.Sleep(100 * time.Millisecond)
			}
		})
	}
}

func TestBookAnalyticsTriggerInvalidPath(t *testing.T) {
	// Initialize the dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	// Create request with invalid path
	req := httptest.NewRequest(http.MethodPost, "/invalid", nil)
	w := httptest.NewRecorder()

	// Call the handler
	h.BookAnalyticsTrigger(w, req)

	// Check status code
	if w.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestBookAnalyticsTriggerResponseHeaders(t *testing.T) {
	// Initialize the dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	// Create request
	req := httptest.NewRequest(http.MethodPost, "/analytics", nil)
	w := httptest.NewRecorder()

	// Call the handler
	h.BookAnalyticsTrigger(w, req)

	// Check content type header
	expectedContentType := "application/json"
	if contentType := w.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("expected Content-Type %s, got %s", expectedContentType, contentType)
	}
}
