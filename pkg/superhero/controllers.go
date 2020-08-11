package superhero

import (
	"encoding/json"
	"net/http"
)

type SuperHeroCtrl interface {
	AddSuper(w http.ResponseWriter, r *http.Request)
	RemoveSuper(w http.ResponseWriter, r *http.Request)
	ListSuper(w http.ResponseWriter, r *http.Request)
	FindSuper(w http.ResponseWriter, r *http.Request)
}

type superHeroCtrlImpl struct {
	srv SuperHeroService
}

func NewSuperHeroCtrl(srv SuperHeroService) SuperHeroCtrl {
	return &superHeroCtrlImpl{srv}
}

func (ctrl *superHeroCtrlImpl) AddSuper(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	err := json.NewDecoder(r.Body).Decode(&params)
	enc := json.NewEncoder(w)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		enc.Encode(&map[string]string{"message": "Could not process request"})
		return
	}

	err = ctrl.srv.AddSuper(params["super_name"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		enc.Encode(&map[string]string{"message": "Could not save super"})
		return
	}

	w.WriteHeader(200)
}

func (ctrl *superHeroCtrlImpl) RemoveSuper(w http.ResponseWriter, r *http.Request) {

}

func (ctrl *superHeroCtrlImpl) ListSuper(w http.ResponseWriter, r *http.Request) {
	superType := r.URL.Query().Get("type")
	resp, err := ctrl.srv.ListSuper(SuperType(superType))

	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(400)
		enc.Encode(&map[string]string{"message": "Could not find any super"})
		return
	}

	err = enc.Encode(resp)

	if err != nil {
		w.WriteHeader(500)
		enc.Encode(&map[string]string{"message": "Could not deliver response"})
		return
	}
}

func (ctrl *superHeroCtrlImpl) FindSuper(w http.ResponseWriter, r *http.Request) {

}
