package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"isbn/dto"
)

// @Summary Delete a book by ID
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} dto.BookDeleteByIDResponse
// @Router /books/{id} [delete]
func (h *Handler) BookDeleteByID(w http.ResponseWriter, r *http.Request) {
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

	var res dto.BookDeleteByIDResponse
	var err error

	res, err = h.service.BookDeleteByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
