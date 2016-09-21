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
type OutResourceBodyHandler struct {
	storage storage.Driver
}

// NewOutResourceBodyHandler creates a new, configured OutResourceBodyHandler.
func NewOutResourceBodyHandler(storage storage.Driver) *OutResourceBodyHandler {
	return &OutResourceBodyHandler{
		storage: storage,
	}
}

// HandlerFunc is the actual HTTP handler.
func (h *OutResourceBodyHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
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

	contentType := request.Headers.Get("Content-Type")
	if contentType != "" {
		w.Header().Add("Content-Type", contentType)
	}

	fmt.Fprint(w, request.Body)
}
