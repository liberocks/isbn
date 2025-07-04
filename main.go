package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"isbn/config"
	"isbn/handler"
	"isbn/logger"
	"isbn/repository"
	"isbn/service"
)

func main() {
	// Initialization
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	// Handle root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to the ISBN API")
	})

	// Handle book-related routes
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.BookGetList(w, r)
			return
		} else if r.Method == http.MethodPost {
			handler.BookCreate(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// Handle book-related routes with ID
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.BookGetByID(w, r)
			return
		} else if r.Method == http.MethodPut {
			handler.BookUpdateByID(w, r)
			return
		} else if r.Method == http.MethodDelete {
			handler.BookDeleteByID(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// Handle book analytics
	http.HandleFunc("/analytics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.BookAnalyticsGet(w, r)
			return
		} else if r.Method == http.MethodPost {
			handler.BookAnalyticsTrigger(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	logger.Logger.Info(fmt.Sprintf("Starting server on :%d", config.AppConfig.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.Port), nil)
}
