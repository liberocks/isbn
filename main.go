package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"isbn/config"
	"isbn/handler"
	"isbn/repository"
	"isbn/service"
)

func main() {
	// Setup logging
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	// Initialization
	repo := repository.NewBookRepository()
	service := service.NewBookService(repo)
	handler := handler.NewBookHandler(service)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to the ISBN API")
	})

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetBooks(w, r)
			return
		} else if r.Method == http.MethodPost {
			handler.CreateBook(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetBookByID(w, r)
			return
		} else if r.Method == http.MethodPut {
			handler.UpdateBookByID(w, r)
			return
		} else if r.Method == http.MethodDelete {
			handler.DeleteBookByID(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	slog.Info(fmt.Sprintf("Starting server on :%d", config.AppConfig.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.Port), nil)
}
