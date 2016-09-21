package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"fmt"

	"github.com/SeerUK/reverb/model"
	"github.com/SeerUK/reverb/storage"
)

// OutCollectionHandler returns the currently stored collection of requests.
type OutCollectionHandler struct {
	storage storage.Driver
}

// NewOutCollectionHandler creates a new, configured OutCollectionHandler.
func NewOutCollectionHandler(storage storage.Driver) *OutCollectionHandler {
	return &OutCollectionHandler{
		storage: storage,
	}
}

// HandlerFunc is the actual HTTP handler.
func (h *OutCollectionHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	var requests []model.Request

	err := h.storage.FindAll(&requests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var collectionItems []OutCollectionResponseItem

	for _, request := range requests {
		collectionItems = append(collectionItems, OutCollectionResponseItem{
			ID:     request.ID,
			Path:   request.Path,
			Method: request.Method,
			Time:   request.Time,
			Link:   fmt.Sprintf("/out/%d", request.ID),
		})
	}

	response := OutCollectionResponse{
		Requests: collectionItems,
	}

	json, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// OutCollectionResponse represents the response for the OutCollectionHandler.
type OutCollectionResponse struct {
	Requests []OutCollectionResponseItem `json:"requests"`
}

// OutCollectionResponseItem represents a simplified version of the model.Request struct, this
// is intended to make output smaller and easier to parse when looking at the collection.
type OutCollectionResponseItem struct {
	ID     uint      `json:"id"`
	Path   string    `json:"path"`
	Method string    `json:"method"`
	Time   time.Time `json:"time"`
	Link   string    `json:"link"`
}
