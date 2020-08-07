package superhero

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewHeroRouter(ctrl SuperHeroCtrl) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome)
	r.HandleFunc("/super", ctrl.AddSuper).Methods("POST")
	r.HandleFunc("/super", ctrl.ListSuper).Methods("GET")
	return cors.Default().Handler(r)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Welcome to superhero API")
	w.Write(msg)
}
