package handler

import (
	"net/http"

	"encoding/json"

	"time"

	"github.com/SeerUK/reverb/model"
	"github.com/SeerUK/reverb/storage"
)

// ReverbCollectionHandler returns the currently stored collection of requests.
type ReverbCollectionHandler struct {
	storage storage.Driver
}

// NewReverbCollectionHandler creates a new, configured ReverbCollectionHandler.
func NewReverbCollectionHandler(storage storage.Driver) *ReverbCollectionHandler {
	return &ReverbCollectionHandler{
		storage: storage,
	}
}

// HandlerFunc is the actual HTTP handler.
func (h *ReverbCollectionHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	var requests []model.Request

	err := h.storage.FindAll(&requests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var collectionItems []ReverbCollectionResponseItem

	for _, request := range requests {
		collectionItems = append(collectionItems, ReverbCollectionResponseItem{
			ID:     request.ID,
			Path:   request.Path,
			Method: request.Method,
			Time:   request.Time,
		})
	}

	response := ReverbCollectionResponse{
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

// ReverbCollectionResponse represents the response for the ReverbCollectionHandler.
type ReverbCollectionResponse struct {
	Requests []ReverbCollectionResponseItem `json:"requests"`
}

// ReverbCollectionResponseItem represents a simplified version of the model.Request struct, this
// is intended to make output smaller and easier to parse when looking at the collection.
type ReverbCollectionResponseItem struct {
	ID     uint      `json:"id"`
	Path   string    `json:"path"`
	Method string    `json:"method"`
	Time   time.Time `json:"time"`
}
