package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"isbn/dto"
)

// @Summary Get a book by ID
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} dto.BookGetByIDResponse
// @Router /books/{id} [get]
func (h *Handler) BookGetByID(w http.ResponseWriter, r *http.Request) {
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

	var res dto.BookGetByIDResponse
	var err error

	res, err = h.service.BookGetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
