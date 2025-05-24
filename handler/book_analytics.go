package handler

import "net/http"

type BookAnalyticsHandlerInterface interface {
	BookAnalyticsTrigger(w http.ResponseWriter, r *http.Request)
	BookAnalyticsGet(w http.ResponseWriter, r *http.Request)
}
