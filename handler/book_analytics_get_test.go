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

func TestBookAnalyticsGet(t *testing.T) {
	// Initialize the dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	tests := []struct {
		name           string
		expectedStatus int
		setupFunc      func()
		cleanupFunc    func()
		validateFunc   func(*dto.BookAnalyticsGetResponse) error
	}{
		{
			name:           "Valid analytics request with data",
			expectedStatus: http.StatusOK,
			setupFunc: func() {
				// Add some test books to generate analytics
				repo.BookCreate(context.Background(), dto.BookCreateRequest{
					ISBN:        "1111111111111",
					Title:       "Test Book 1",
					Author:      "Author 1",
					ReleaseDate: "2023-01-01",
				})
				repo.BookCreate(context.Background(), dto.BookCreateRequest{
					ISBN:        "1111111111112",
					Title:       "Test Book 1 Sequel",
					Author:      "Author 1",
					ReleaseDate: "2023-01-01",
				})
				repo.BookCreate(context.Background(), dto.BookCreateRequest{
					ISBN:        "2222222222222",
					Title:       "Test Book 2A",
					Author:      "Author 2",
					ReleaseDate: "2023-06-01",
				})

				// Trigger analytics calculation
				triggerReq := httptest.NewRequest(http.MethodPost, "/analytics", nil)
				triggerW := httptest.NewRecorder()
				h.BookAnalyticsTrigger(triggerW, triggerReq)
			},
			cleanupFunc: func() {
				repo.BookDeleteByID(context.Background(), "1111111111111")
				repo.BookDeleteByID(context.Background(), "1111111111112")
				repo.BookDeleteByID(context.Background(), "2222222222222")
			},
			validateFunc: func(response *dto.BookAnalyticsGetResponse) error {
				return nil
			},
		},
		{
			name:           "Empty repository",
			expectedStatus: http.StatusOK,
			validateFunc: func(response *dto.BookAnalyticsGetResponse) error {
				// For empty repository, all values should be zero/empty
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				tt.setupFunc()
			}

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/analytics", nil)

			// Create response recorder
			w := httptest.NewRecorder()

			// Call the handler
			h.BookAnalyticsGet(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d. Response body: %s", tt.expectedStatus, w.Code, w.Body.String())
			}

			// For successful requests, verify response structure
			if tt.expectedStatus == http.StatusOK {
				var response dto.BookAnalyticsGetResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Fatalf("failed to decode response: %v. Response body: %s", err, w.Body.String())
				}

				// Validate response using custom validation function
				if tt.validateFunc != nil {
					if err := tt.validateFunc(&response); err != nil {
						t.Errorf("validation failed: %v", err)
					}
				}

				t.Logf("Analytics response: TotalBooks=%d, TotalAuthors=%d", response.TotalBooks, response.TotalAuthors)
			}

			// Cleanup
			if tt.cleanupFunc != nil {
				tt.cleanupFunc()
			}
		})
	}
}

func TestBookAnalyticsGetInvalidPath(t *testing.T) {
	// Initialize the dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	// Create request with invalid path
	req := httptest.NewRequest(http.MethodGet, "/invalid", nil)
	w := httptest.NewRecorder()

	// Call the handler
	h.BookAnalyticsGet(w, req)

	// Check status code
	if w.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}
