package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fmt"

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

	// Set the body to be a link to the URL to get the body. We don't want to clog up the output
	// JSON here with what could be anything really.
	request.Body = fmt.Sprintf("/out/%d/body", request.ID)

	json, err := json.Marshal(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
