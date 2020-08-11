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
	params := mux.Vars(r)
	uuid, err := strconv.ParseInt(params["uuid"], 10, 64)

	if err != nil {
		writeRequestError(w, 500, "Could not process request")
		return
	}

	resp, err := ctrl.srv.FindByUUID(uuid)
	if err != nil {
		writeRequestError(w, 500, "Could not process request")
		return
	}

	err = writeRequestResponse(w, resp)

	if err != nil {
		writeRequestError(w, 500, "Could not deliver response")
	}
}

func (ctrl *ctrlImpl) FindByNameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if name, ok := params["name"]; ok {
		resp, err := ctrl.srv.FindByName(name)

		if err != nil {
			writeRequestError(w, 500, "Could not process request")
			return
		}

		err = writeRequestResponse(w, resp)

		if err != nil {
			writeRequestError(w, 500, "Could not deliver response")
		}
	} else {
		writeRequestError(w, 500, "Could not process request")
	}
}

func (ctrl *ctrlImpl) ListHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var superType SuperType
	if t, ok := params["type"]; ok {
		superType = SuperType(t)
	} else {
		superType = NoType
	}

	resp, err := ctrl.srv.ListByType(superType)
	if err != nil {
		writeRequestError(w, 500, "Could not process request")
		return
	}

	err = writeRequestResponse(w, resp)

	if err != nil {
		writeRequestError(w, 500, "Could not deliver response")
	}
}

func writeRequestResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func writeRequestError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorMsg := newErrorResult(msg)
	json.NewEncoder(w).Encode(errorMsg)
}
