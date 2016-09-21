package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/SeerUK/reverb/model"
	"github.com/SeerUK/reverb/storage"
)

// ApiHandlerResponse represents the structure of the response given by the InHandler.
type InHandlerResponse struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
}

// ApiHandler records all requests that come to it in some storage.
type InHandler struct {
	storage storage.Driver
}

// NewInHandler creates a new, configured InHandler.
func NewInHandler(storage storage.Driver) *InHandler {
	return &InHandler{
		storage: storage,
	}
}

// HandlerFunc is the actual HTTP handler.
func (h *InHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	request := model.Request{
		Path:    strings.TrimPrefix(r.RequestURI, "/api"),
		Method:  r.Method,
		Headers: r.Header,
		Time:    time.Now(),
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	request.Body = string(body)

	h.storage.Persist(&request)

	json, err := json.Marshal(InHandlerResponse{
		ID:      request.ID,
		Message: "OK",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
