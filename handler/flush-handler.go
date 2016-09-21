package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SeerUK/reverb/storage"
)

// FlushHandlerResponse represents the structure of the response given by the FlushHandler.
type FlushHandlerResponse struct {
	Message string `json:"message"`
}

// FlushHandler records all requests that come to it in some storage.
type FlushHandler struct {
	storage storage.Driver
}

// NewFlushHandler creates a new, configured FlushHandler.
func NewFlushHandler(storage storage.Driver) *FlushHandler {
	return &FlushHandler{
		storage: storage,
	}
}

// HandlerFunc is the actual HTTP handler.
func (h *FlushHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	h.storage.Flush()

	json, err := json.Marshal(FlushHandlerResponse{
		Message: "OK",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
