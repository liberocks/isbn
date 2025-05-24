package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"isbn/dto"
	"isbn/logger"
)

// @Summary Update a book by ID
// @Produce json
// @Param id path string true "Book ID"
// @Param book body dto.BookUpdateRequest true "Book object"
// @Success 200 {object} dto.BookUpdateByIDResponse
func (h *Handler) BookUpdateByID(w http.ResponseWriter, r *http.Request) {
	prefix := "/books/"

	if !strings.HasPrefix(r.URL.Path, prefix) {
		logger.Logger.Error("Invalid URL path", "path", r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, prefix)
	if id == "" {
		logger.Logger.Error("Book ID is required", "path", r.URL.Path)
		http.Error(w, "Book ID is required", http.StatusBadRequest)
		return
	}

	// Parse the request body
	var req dto.BookUpdateByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Logger.Error("Failed to decode request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		logger.Logger.Error("Validation error", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the book
	var res dto.BookUpdateByIDResponse
	var err error

	res, err = h.service.BookUpdateByID(r.Context(), id, req)
	if err != nil {
		logger.Logger.Error("Failed to update book", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
