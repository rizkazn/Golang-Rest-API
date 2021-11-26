package controllers

import "net/http"

type Controls interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	SearchByName(w http.ResponseWriter, r *http.Request)
	SortByName(w http.ResponseWriter, r *http.Request)
	SortByCat(w http.ResponseWriter, r *http.Request)
	SortByNewest(w http.ResponseWriter, r *http.Request)
	SortByPrice(w http.ResponseWriter, r *http.Request)
}
