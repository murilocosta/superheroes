package superhero

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type errorResult struct {
	message string
}

func newErrorResult(msg string) *errorResult {
	return &errorResult{message: msg}
}

type Ctrl interface {
	CreateHandler(w http.ResponseWriter, r *http.Request)
	DeleteHandler(w http.ResponseWriter, r *http.Request)
	FindByUUIDHandler(w http.ResponseWriter, r *http.Request)
	FindByNameHandler(w http.ResponseWriter, r *http.Request)
	ListHandler(w http.ResponseWriter, r *http.Request)
}

type ctrlImpl struct {
	srv Service
}

func NewCtrl(srv Service) Ctrl {
	return &ctrlImpl{srv}
}

func (ctrl *ctrlImpl) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		writeRequestError(w, 500, "Could not process request")
		return
	}

	err = ctrl.srv.Create(params["super_name"])

	if err != nil {
		writeRequestError(w, 400, "Could not save super")
		return
	}

	w.WriteHeader(200)
}

func (ctrl *ctrlImpl) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	superID, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		writeRequestError(w, 500, "Could not process request")
		return
	}

	err = ctrl.srv.Delete(superID)

	if err != nil {
		writeRequestError(w, 400, "Could not delete super")
		return
	}
}

func (ctrl *ctrlImpl) FindByUUIDHandler(w http.ResponseWriter, r *http.Request) {

}

func (ctrl *ctrlImpl) FindByNameHandler(w http.ResponseWriter, r *http.Request) {

}

func (ctrl *ctrlImpl) ListHandler(w http.ResponseWriter, r *http.Request) {
	superType := r.URL.Query().Get("type")
	resp, err := ctrl.srv.ListByType(SuperType(superType))

	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		writeRequestErrorJSON(w, enc, 400, "Could not find any super")
		return
	}

	err = enc.Encode(resp)

	if err != nil {
		writeRequestErrorJSON(w, enc, 500, "Could not deliver response")
		return
	}
}

func writeRequestError(w http.ResponseWriter, status int, msg string) {
	enc := json.NewEncoder(w)
	writeRequestErrorJSON(w, enc, status, msg)
}

func writeRequestErrorJSON(w http.ResponseWriter, enc *json.Encoder, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorMsg := newErrorResult(msg)
	enc.Encode(errorMsg)
}
