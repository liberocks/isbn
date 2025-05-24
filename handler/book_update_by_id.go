package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"isbn/dto"
)

// @Summary Update a book by ID
// @Produce json
// @Param id path string true "Book ID"
// @Param book body dto.BookUpdateRequest true "Book object"
// @Success 200 {object} dto.BookUpdateByIDResponse
func (h *Handler) BookUpdateByID(w http.ResponseWriter, r *http.Request) {
	prefix := "/books/"

	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, prefix)
	if id == "" {
		http.Error(w, "Book ID is required", http.StatusBadRequest)
		return
	}

	// Parse the request body
	var req dto.BookUpdateByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the book
	var res dto.BookUpdateByIDResponse
	var err error

	res, err = h.service.BookUpdateByID(r.Context(), id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
