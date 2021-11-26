package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizkazn/gorestfull/controllers"
	"github.com/rizkazn/gorestfull/repos"
)

func CategoryRoute(r *mux.Router, db *sql.DB) {

	catrepo := repos.Category(db)
	cr := controllers.Category(catrepo)

	// Prefix and Subrouter
	route := r.PathPrefix("/categories").Subrouter()
	// Route Handlers / Endpoints
	route.HandleFunc("", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.Add).Methods(http.MethodPost)
	route.HandleFunc("/{id}", cr.Update).Methods(http.MethodPut)
	route.HandleFunc("/{id}", cr.Delete).Methods(http.MethodDelete)
}
