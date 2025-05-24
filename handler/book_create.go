package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"isbn/dto"
)

// @Summary Create a new book
// @Produce json
// @Param book body dto.BookCreateRequest true "Book object"
// @Success 201 {object} dto.BookCreateResponse
// @Router /books [post]
func (h *Handler) BookCreate(w http.ResponseWriter, r *http.Request) {
	prefix := "/books"

	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Parse the request body
	var req dto.BookCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the book
	res, err := h.service.BookCreate(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
