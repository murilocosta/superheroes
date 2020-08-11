package superhero

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter(ctrl Ctrl) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/supers", ctrl.CreateHandler).Methods("POST")
	r.HandleFunc("/supers", ctrl.ListHandler).Methods("GET")
	r.HandleFunc("/supers/type/{type}", ctrl.ListHandler).Methods("GET")
	r.HandleFunc("/supers/name/{name}", ctrl.FindByNameHandler).Methods("GET")
	r.HandleFunc("/supers/{uuid}", ctrl.FindByUUIDHandler).Methods("GET")
	r.HandleFunc("/supers/{id}", ctrl.DeleteHandler).Methods("DELETE")
	return cors.Default().Handler(r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Welcome to SuperHeroes API")
	w.Write(msg)
}
