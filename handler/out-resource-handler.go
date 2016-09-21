package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SeerUK/reverb/model"
	"github.com/SeerUK/reverb/storage"
	"github.com/gorilla/mux"
)

// OutResourceHandler returns the the details of a stored request.
type OutResourceHandler struct {
	storage storage.Driver
}

// NewOutResourceHandler creates a new, configured OutResourceHandler.
func NewOutResourceHandler(storage storage.Driver) *OutResourceHandler {
	return &OutResourceHandler{
		storage: storage,
	}
}

// HandlerFunc is the actual HTTP handler.
func (h *OutResourceHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	var request model.Request

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.storage.Find(id, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Clear the body, we don't want to clog up the response body here
	request.Body = ""

	json, err := json.Marshal(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
