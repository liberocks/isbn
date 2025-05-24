package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"isbn/dto"
	"isbn/logger"
)

// @Summary Get book analytics
// @Produce json
// @Success 200 {object} dto.BookAnalyticsGetResponse
// @Router /analytics [get]
func (h *Handler) BookAnalyticsGet(w http.ResponseWriter, r *http.Request) {
	prefix := "/analytics"

	if !strings.HasPrefix(r.URL.Path, prefix) {
		logger.Logger.Error("Invalid URL path", "path", r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Respond with 200 OK
	var res dto.BookAnalyticsGetResponse
	var err error

	res, err = h.service.BookAnalyticsGet(r.Context())
	if err != nil {
		logger.Logger.Error("Failed to get book analytics", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the analytics data
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
