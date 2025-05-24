package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"isbn/dto"
	"isbn/logger"
)

// @Summary Delete a book by ID
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} dto.BookDeleteByIDResponse
// @Router /books/{id} [delete]
func (h *Handler) BookDeleteByID(w http.ResponseWriter, r *http.Request) {
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

	var res dto.BookDeleteByIDResponse
	var err error

	res, err = h.service.BookDeleteByID(r.Context(), id)
	if err != nil {
		logger.Logger.Error("Failed to delete book", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
