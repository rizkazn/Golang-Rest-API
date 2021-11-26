package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizkazn/gorestfull/controllers"
	"github.com/rizkazn/gorestfull/repos"
)

func AuthRoute(r *mux.Router, db *sql.DB) {

	repo := repos.User(db)
	cr := controllers.Auth(repo)

	// Prefix and Subrouter
	route := r.PathPrefix("/auth").Subrouter()
	// Route Handlers / Endpoints
	route.HandleFunc("/", cr.Login).Methods(http.MethodPost)
}
