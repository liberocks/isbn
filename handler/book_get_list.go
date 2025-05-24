package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"isbn/dto"
	"isbn/logger"
)

// @Summary Get list of books
// @Produce json
// @Success 200 {object} dto.BookGetListResponse
// @Router /books [get]
func (h *Handler) BookGetList(w http.ResponseWriter, r *http.Request) {
	prefix := "/books"

	if !strings.HasPrefix(r.URL.Path, prefix) {
		logger.Logger.Error("Invalid URL path", "path", r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Parse the query string
	var req dto.BookGetListQuery

	query := r.URL.Query()
	limit := query.Get("limit")
	page := query.Get("page")

	if limit != "" {
		req.Limit, _ = strconv.Atoi(limit)
	} else {
		req.Limit = 10 // Default limit
	}

	if page != "" {
		req.Page, _ = strconv.Atoi(page)
	} else {
		req.Page = 1 // Default page
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		logger.Logger.Error("Validation error", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var res dto.BookGetListResponse
	var err error

	res, err = h.service.BookGetList(r.Context(), req)
	if err != nil {
		logger.Logger.Error("Failed to get book list", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
