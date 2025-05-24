package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"isbn/dto"
	"isbn/logger"
)

// @Summary Trigger book analytics
// @Produce json
// @Success 200 {object} dto.BookAnalyticsTriggerResponse
// @Router /analytics [post]
func (h *Handler) BookAnalyticsTrigger(w http.ResponseWriter, r *http.Request) {
	prefix := "/analytics"

	if !strings.HasPrefix(r.URL.Path, prefix) {
		logger.Logger.Error("Invalid URL path", "path", r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Trigger the book analytics without waiting for the result
	go func() {
		err := h.service.BookAnalyticsTrigger(r.Context())
		if err != nil {
			logger.Logger.Error("Failed to trigger book analytics", "error", err)
			// Log the error but do not block the response
			return
		}
		logger.Logger.Info("Book analytics triggered successfully")
	}()

	// Respond with 200 OK
	var res dto.BookAnalyticsTriggerResponse
	res.Message = "Book analytics triggered successfully"

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
