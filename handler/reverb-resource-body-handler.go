package handler

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/SeerUK/reverb/model"
	"github.com/SeerUK/reverb/storage"
	"github.com/gorilla/mux"
)

// ReverbResourceBodyHandler returns the body of a specific stored request.
type ReverbResourceBodyHandler struct {
	storage storage.Driver
}

// NewReverbResourceBodyHandler creates a new, configured ReverbResourceBodyHandler.
func NewReverbResourceBodyHandler(storage storage.Driver) *ReverbResourceBodyHandler {
	return &ReverbResourceBodyHandler{
		storage: storage,
	}
}

// HandlerFunc is the actual HTTP handler.
func (h *ReverbResourceBodyHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprint(w, request.Body)
}
