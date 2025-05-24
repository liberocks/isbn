package handler

import "net/http"

type BookHandlerInterface interface {
	BookCreate(w http.ResponseWriter, r *http.Request)
	BookGetList(w http.ResponseWriter, r *http.Request)
	BookGetByID(w http.ResponseWriter, r *http.Request)
	BookUpdateByID(w http.ResponseWriter, r *http.Request)
	BookDeleteByID(w http.ResponseWriter, r *http.Request)
}
