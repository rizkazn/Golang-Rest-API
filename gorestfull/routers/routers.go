package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizkazn/gorestfull/configs/db"
)

func New() *mux.Router {
	r := mux.NewRouter()

	// inisialisasi endpoint
	r.HandleFunc("/", simpleHandlers).Methods(http.MethodGet)
	mainRutes := r.PathPrefix("/api/v1").Subrouter().StrictSlash(false)
	// mainRutes.HandleFunc("/foo", fooHandlers).Methods(http.MethodGet, http.MethodOptions)

	// inisialisasi dbms
	dbms, _ := db.New()

	mainRutes.Use(mux.CORSMethodMiddleware(mainRutes))

	ProductRoute(mainRutes, dbms)
	CategoryRoute(mainRutes, dbms)
	HistoryRoute(mainRutes, dbms)
	UserRoute(mainRutes, dbms)
	AuthRoute(mainRutes, dbms)

	return mainRutes

}

func simpleHandlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API"))
}

// func fooHandlers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access.Control.Allow.Origin", "*")
// 	w.Write([]byte("Hello from Foo"))
// }
